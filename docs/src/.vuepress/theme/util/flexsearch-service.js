import _ from "lodash";
import Flexsearch from "flexsearch";

let pagesByPath = null;
let indexes = [];

export default {
  buildIndex(pages) {
    const indexSettings = {
      async: true,
      doc: {
        id: "key",
        // fields to index
        field: ["title", "keywords", "headersStr", "content"]
      }
    };

    const globalIndex = new Flexsearch(indexSettings);
    globalIndex.add(pages);

    indexes["global"] = globalIndex;

    pagesByPath = _.keyBy(pages, "path");
  },

  async match(queryString, queryTerms, limit = 7) {
    const index = resolveSearchIndex(indexes);

    const searchParams = [
      {
        field: "keywords",
        query: queryString,
        boost: 8,
        suggest: false,
        bool: "or",
      },
      {
        field: "title",
        query: queryString,
        boost: 10,
        suggest: false,
        bool: "or",
      },
      {
        field: "headersStr",
        query: queryString,
        boost: 7,
        suggest: false,
        bool: "or",
      },
      {
        field: "content",
        query: queryString,
        boost: 0,
        suggest: false,
        bool: "or",
      }
    ];

    const searchResult = await index.search(searchParams, limit);
    const result = searchResult.map(page => ({
      ...page,
      parentPageTitle: getParentPageTitle(page),
      ...getAdditionalInfo(page, queryString, queryTerms)
    }));

    return result.map((x, i) => {
      if (i === 0) return x;
      return { ...x, parentPageTitle: null };
    });
  }
};

function resolveSearchIndex(indexes) {
  return indexes["global"];
}

function getParentPageTitle(page) {
  const pathParts = page.path.split("/");
  let parentPagePath = "/";
  if (pathParts[1]) parentPagePath = `/${pathParts[1]}/`;

  const parentPage = pagesByPath[parentPagePath] || page;
  return parentPage.title;
}

/**
 * Returns contextual details for displaying search result.
 * @param {*} page
 * @param {*} queryString
 * @param {*} queryTerms
 */
function getAdditionalInfo(page, queryString, queryTerms) {
  const query = queryString.toLowerCase();

  /**
   * If it’s an exact title match or the page title starts with the query string,
   * return the result with the full heading and no slug.
   */
  if (
    page.title.toLowerCase() === query ||
    page.title.toLowerCase().startsWith(query)
  ) {
    return {
      headingStr: getFullHeading(page),
      slug: "",
      contentStr: getBeginningContent(page),
      match: "title"
    };
  }

  /**
   * If our special (and pretty much invisible) keywords include the query string,
   * return the result using the page title, no slug, and opening sentence.
   */
  if (page.keywords.includes(query)) {
    return {
      headingStr: getFullHeading(page),
      slug: "",
      contentStr: getBeginningContent(page),
      match: "keywords"
    };
  }

  const match = getMatch(page, query, queryTerms);

  /**
   * If we can’t match the query string to anything specific, list the result
   * with only the page heading.
   */
  if (!match)
    return {
      headingStr: getFullHeading(page),
      slug: "",
      contentStr: null,
      match: "?"
    };

  /**
   * If we have a match that’s in a heading, display that heading and return
   * a link to it without any content snippet.
   */
  if (match.headerIndex != null) {
    // header match
    return {
      headingStr: getFullHeading(page, match.headerIndex),
      slug: "#" + page.headers[match.headerIndex].slug,
      contentStr: null,
      match: "header"
    };
  }

  /**
   * Get the index of the nearest preceding header relative to the content match.
   */
  let headerIndex = _.findLastIndex(
    page.headers || [],
    h => h.charIndex != null && h.charIndex < match.charIndex
  );
  if (headerIndex === -1) headerIndex = null;

  return {
    headingStr: getFullHeading(page, headerIndex),
    slug: headerIndex == null ? "" : "#" + page.headers[headerIndex].slug,
    contentStr: getContentStr(page, match),
    match: "content"
  };
}

/**
 * Return the target heading in the context of its parents. (Like a breadcrumb.)
 * @param {*} page
 * @param {*} headerIndex
 */
function getFullHeading(page, headerIndex) {
  if (headerIndex == null) return page.title;
  const headersPath = [];
  while (headerIndex != null) {
    const header = page.headers[headerIndex];
    headersPath.unshift(header);
    headerIndex = _.findLastIndex(
      page.headers,
      h => h.level === header.level - 1,
      headerIndex - 1
    );
    if (headerIndex === -1) headerIndex = null;
  }
  return headersPath.map(h => h.title).join(" → ");
}

function getMatch(page, query, terms) {
  const matches = terms
    .map(t => {
      return getHeaderMatch(page, t) || getContentMatch(page, t);
    })
    .filter(m => m);
  if (matches.length === 0) return null;

  if (matches.every(m => m.headerIndex != null)) {
    return getHeaderMatch(page, query) || matches[0];
  }

  return (
    getContentMatch(page, query) || matches.find(m => m.headerIndex == null)
  );
}

function getHeaderMatch(page, term) {
  if (!page.headers) return null;
  for (let i = 0; i < page.headers.length; i++) {
    const h = page.headers[i];
    const charIndex = h.title.toLowerCase().indexOf(term);
    if (charIndex === -1) continue;
    return {
      headerIndex: i,
      charIndex,
      termLength: term.length
    };
  }
  return null;
}

function getContentMatch(page, term) {
  if (!page.contentLowercase) return null;
  const charIndex = page.contentLowercase.indexOf(term);
  if (charIndex === -1) return null;

  return { headerIndex: null, charIndex, termLength: term.length };
}

function getContentStr(page, match) {
  const snippetLength = 120;
  const { charIndex, termLength } = match;

  let lineStartIndex = page.content.lastIndexOf("\n", charIndex);
  let lineEndIndex = page.content.indexOf("\n", charIndex);

  if (lineStartIndex === -1) lineStartIndex = 0;
  if (lineEndIndex === -1) lineEndIndex = page.content.length;

  const line = page.content.slice(lineStartIndex, lineEndIndex);

  if (snippetLength >= line.length) return line;

  const lineCharIndex = charIndex - lineStartIndex;

  const additionalCharactersFromStart = (snippetLength - termLength) / 2;
  const snippetStart = Math.max(
    lineCharIndex - additionalCharactersFromStart,
    0
  );
  const snippetEnd = Math.min(snippetStart + snippetLength, line.length);
  let result = line.slice(snippetStart, snippetEnd);
  if (snippetStart > 0) result = "..." + result;
  if (snippetEnd < line.length) result = result + "...";
  return result;
}

/**
 * Returns the initial page content after the title.
 * @param {*} page
 */
function getBeginningContent(page) {
  const lines = page.contentLowercase.split("\n");
  const lowerFirstLine = (lines.length > 0 ? lines[0] : "").trim();
  const lowerPageTitle = page.title.toLowerCase()
  // the first line is the title, or the title with an edition badge
  const firstLineIsTitle = lowerFirstLine === lowerPageTitle ||
    lowerFirstLine === `${lowerPageTitle} pro` ||
    lowerFirstLine === `${lowerPageTitle} lite` ||
    lowerFirstLine === `${lowerPageTitle} solo`

  if (firstLineIsTitle) {
    // first line *is* title; start at second line
    return getContentStr(page, {
      charIndex: lowerFirstLine.length + 2,
      termLength: 0
    });
  }

  return getContentStr(page, { charIndex: 0, termLength: 0 });
}
