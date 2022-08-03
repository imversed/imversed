export const hashRE = /#.*$/;
export const extRE = /\.(md|html)$/;
export const endingSlashRE = /\/$/;
export const outboundRE = /^[a-z]+:/i;

export function normalize(path) {
  return decodeURI(path)
    .replace(hashRE, "")
    .replace(extRE, "");
}

export function getHash(path) {
  const match = path.match(hashRE);
  if (match) {
    return match[0];
  }
}

export function isExternal(path) {
  return outboundRE.test(path);
}

export function isMailto(path) {
  return /^mailto:/.test(path);
}

export function isTel(path) {
  return /^tel:/.test(path);
}

export function ensureExt(path) {
  if (isExternal(path)) {
    return path;
  }
  const hashMatch = path.match(hashRE);
  const hash = hashMatch ? hashMatch[0] : "";
  const normalized = normalize(path);

  if (endingSlashRE.test(normalized)) {
    return path;
  }
  return normalized + ".html" + hash;
}

export function isActive(route, path) {
  const routeHash = decodeURIComponent(route.hash);
  const linkHash = getHash(path);
  if (linkHash && routeHash !== linkHash) {
    return false;
  }
  const routePath = normalize(route.path);
  const pagePath = normalize(path);
  return routePath === pagePath;
}

export function resolvePage(pages, rawPath, base) {
  if (isExternal(rawPath)) {
    return {
      type: "external",
      path: rawPath
    };
  }
  if (base) {
    rawPath = resolvePath(rawPath, base);
  }
  const path = normalize(rawPath);
  for (let i = 0; i < pages.length; i++) {
    if (normalize(pages[i].regularPath) === path) {
      const resolved = Object.assign({}, pages[i], {
        type: "page",
        path: ensureExt(pages[i].path)
      });
      return resolved;
    }
  }
  console.error(
    `[vuepress] No matching page found for sidebar item "${rawPath}"`
  );
  return {};
}

function resolvePath(relative, base, append) {
  const firstChar = relative.charAt(0);
  if (firstChar === "/") {
    return relative;
  }

  if (firstChar === "?" || firstChar === "#") {
    return base + relative;
  }

  const stack = base.split("/");

  // remove trailing segment if:
  // - not appending
  // - appending to trailing slash (last segment is empty)
  if (!append || !stack[stack.length - 1]) {
    stack.pop();
  }

  // resolve relative path
  const segments = relative.replace(/^\//, "").split("/");
  for (let i = 0; i < segments.length; i++) {
    const segment = segments[i];
    if (segment === "..") {
      stack.pop();
    } else if (segment !== ".") {
      stack.push(segment);
    }
  }

  // ensure leading slash
  if (stack[0] !== "") {
    stack.unshift("");
  }

  return stack.join("/");
}

/**
 * @param { Page } page
 * @param { string } regularPath
 * @param { SiteData } site
 * @param { string } localePath
 * @returns { SidebarGroup }
 */
export function resolveSidebarItems(
  page,
  regularPath,
  site,
) {
  const { pages, themeConfig } = site;

  // get the config object for whatever sidebar items we should be showing
  const sidebarConfig = resolveSidebarConfig(page, themeConfig);

  if (!sidebarConfig) {
    return [];
  } else {
    // get the correct sidebar, whether the config is an array or path-indexed object
    let { base, config } = resolveMatchingConfig(regularPath, sidebarConfig);

    if (!config) {
      console.log("didn’t resolve config");
      return [];
    }

    const resolved = config.map(item => {
      return resolveItem(item, pages, base);
    });

    return resolved;
  }
}

/**
 * @param { Page } page
 * @param { string } regularPath
 * @param { SiteData } site
 * @returns { SidebarGroup }
 */
export function resolveExtraSidebarItems(
  page,
  regularPath,
  site,
) {
  const { pages, themeConfig } = site;

  // get the config object for whatever sidebar items we should be showing
  const sidebarConfig = resolveExtraSidebarConfig(themeConfig);

  if (!sidebarConfig) {
    return [];
  } else {
    // get the correct sidebar, whether the config is an array or path-indexed object
    let { config } = resolveMatchingConfig(regularPath, sidebarConfig);

    if (!config) {
      console.log("didn’t resolve config");
      return [];
    }

    const resolved = config.map(item => {
      return resolveExtraItem(item);
    });

    return resolved;
  }
}

export function resolveSidebarConfig(
  page,
  themeConfig
) {
  let sidebarConfig;

  if (page.frontmatter.sidebar) {
    sidebarConfig = page.frontmatter.sidebar;
  } else if (themeConfig.sidebar) {
    sidebarConfig = themeConfig.sidebar;
  }

  return sidebarConfig;
}

export function resolveExtraSidebarConfig(themeConfig) {
  return themeConfig.sidebarExtra || [];
}

/**
 * Translate page content headers into sidebar items.
 * @param { Page } page
 * @returns { SidebarGroup }
 */
export function resolveHeaders(page) {
  const headers = groupHeaders(
    page.headers || [],
    page.frontmatter.sidebarLevel
  );
  return [
    {
      type: "group",
      collapsable: false,
      title: page.title,
      path: null,
      children: headers.map(h => ({
        type: "auto",
        title: h.title,
        basePath: page.path,
        path: page.path + "#" + h.slug,
        children: h.children || []
      }))
    }
  ];
}

/**
 * Collect headers grouped by specified target level. (Default is `h2`.)
 * @param {*} headers
 * @param {*} level
 */
export function groupHeaders(headers, level = 2) {
  // normalize objects
  headers = headers.map(h => Object.assign({}, h));
  let lastHeadingAtLevel;

  // collect children of target level
  headers.forEach(h => {
    if (h.level === level) {
      lastHeadingAtLevel = h;
    } else if (lastHeadingAtLevel) {
      (lastHeadingAtLevel.children || (lastHeadingAtLevel.children = [])).push(
        h
      );
    }
  });

  return headers.filter(h => h.level === level);
}

export function resolveNavLinkItem(linkItem) {
  return Object.assign(linkItem, {
    type: linkItem.items && linkItem.items.length ? "links" : "link"
  });
}

/**
 * Takes the regular path (like `/3.x/extend/widget-types.html`) and locale-resolved config
 * to return the current base and relevant section of the sidebar config.
 *
 * Modified to account for the active docSet and version.
 *
 * @param { Route } route
 * @param { Array<string|string[]> | Array<SidebarGroup> | [link: string]: SidebarConfig } config
 * @returns { base: string, config: SidebarConfig }
 */
export function resolveMatchingConfig(regularPath, config) {
  // always starts with `/`
  let base = "/";

  base = fixDoubleSlashes(ensureEndingSlash(base));

  // simpler array
  if (Array.isArray(config)) {
    return {
      base: base,
      config: config
    };
  }

  // sidebar config by path, where `/` is the default
  const activeBase = getRelativeActiveBaseFromConfig(
    regularPath,
    config
  );

  if (activeBase) {
    return {
      base: fixDoubleSlashes(base + activeBase),
      config: config[activeBase]
    };
  }

  return {};
}

/**
 * Returns the active sidebar config key.
 *
 * @param {*} path
 * @param {*} config
 */
export function getRelativeActiveBaseFromConfig(path, config) {
  if (Array.isArray(config)) {
    return;
  }

  for (const activeBase in config) {
    if (ensureEndingSlash(path).indexOf(encodeURI(activeBase)) === 0) {
      return activeBase;
    }
  }

  return;
}

function ensureEndingSlash(path) {
  return /(\.html|\/)$/.test(path) ? path : path + "/";
}

export function fixDoubleSlashes(path) {
  return path.replace(/\/\//g, "/");
}

/**
 * Find the given item among the available pages, taking into account
 * the provided base and depth.
 *
 * @param {*} item  Can be a string like `coc`, or an object with `title`,`collapsable` and `children`.
 * @param {*} pages
 * @param {*} base
 * @param {*} groupDepth
 */
export function resolveItem(item, pages, base, groupDepth = 1) {
  if (typeof item === "string") {
    return resolvePage(pages, item, base);
  } else if (Array.isArray(item)) {
    return Object.assign(resolvePage(pages, item[0], base), {
      title: item[1]
    });
  } else {
    const children = item.children || [];
    if (children.length === 0 && item.path) {
      return Object.assign(resolvePage(pages, item.path, base), {
        title: item.title
      });
    }
    const toggleChildren = item.toggleChildren || [];
    return {
      type: "group",
      path: item.path,
      title: item.title,
      sidebarDepth: item.sidebarDepth,
      children: children.map(child =>
        resolveItem(child, pages, base, groupDepth + 1)
      ),
      toggleChildren: toggleChildren.map(child =>
        resolveItem(child, pages, base, groupDepth + 1)
      ),
      collapsable: item.collapsable !== false
    };
  }
}

/**
 * Find the given item among the available pages, taking into account
 * the provided base and depth.
 *
 * @param {*} item  Can be a string like `coc`, or an object with `title`,`collapsable` and `children`.
 * @param {*} pages
 * @param {*} base
 * @param {*} groupDepth
 */
export function resolveExtraItem(item) {
  return {
    path: item.path,
    title: item.title,
    link: item.link,
    icon: item.icon,
    sidebarDepth: item.sidebarDepth
  };
}
