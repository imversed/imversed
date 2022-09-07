module.exports = {
  base: "/",
  plugins: [
    ["@vuepress/google-analytics", { ga: "UA-39036834-9" }],
    [
      "vuepress-plugin-medium-zoom",
      {
        selector: ".theme-default-content img:not(.no-zoom)",
        delay: 1000,
        options: {
          margin: 24,
          background: "var(--medium-zoom-overlay-color)",
          scrollOffset: 0
        }
      }
    ],
    ["vuepress-plugin-container", { type: "tip", defaultTitle: "" }],
    ["vuepress-plugin-container", { type: "warning", defaultTitle: "" }],
    ["vuepress-plugin-container", { type: "danger", defaultTitle: "" }],
    [
      "vuepress-plugin-container",
      {
        type: "details",
        before: info =>
          `<details class="custom-block details">${
            info ? `<summary>${info}</summary>` : ""
          }\n`,
        after: () => "</details>\n"
      }
    ]
  ],
  shouldPrefetch: () => false,
  head: [
    [
      "link",
      {
        rel: "icon",
        type: "image/png",
        sizes: "32x32",
        href: "https://imversed.com/assets/images/imv_logo.png"
      }
    ],
    ["meta", { name: "msapplication-TileColor", content: "#f1f5fd" }],
    ["meta", { name: "theme-color", content: "#1a202c", media: "(prefers-color-scheme: dark)" }],
    ["meta", { name: "theme-color", content: "#f1f5fd" }]
  ],
  themeConfig: {
    title: "Imversed Documentation",
    docsRepo: "craftcms/docs",
    docsDir: "docs",
    docsBranch: "main",
    baseUrl: "https://docs.imversed.com",
    nextLinks: true,
    prevLinks: true,
    searchMaxSuggestions: 10,
    nav: [
      { text: "Knowlege Base", link: "https://craftcms.com/knowledge-base" }
    ],
    codeLanguages: {
      twig: "Twig",
      php: "PHP",
      graphql: "GraphQL",
      js: "JavaScript",
      json: "JSON",
      xml: "XML",
      treeview: "Folder",
      graphql: "GraphQL",
      csv: "CSV"
    },
    feedback: {
      helpful: "Was this page helpful?",
      thanks: "Thanks for your feedback.",
      more: "Give More Feedback →"
    },
    searchPlaceholder: "Search Imversed Docs (Press “/” to focus)",
    sidebar: {
      "/": [
        {
          title: "Introduction",
          collapsable: false,
          children: ["introduction/", "introduction/resources", "modules/"],
        },
        {
          title: "For Users",
          collapsable: false,
          children: ["forusers/", "forusers/digitalwallets", "forusers/accountkeys", "forusers/imversedgovernance", "forusers/technicalconcepts"],
        },
        {
          title: "For Developers",
          collapsable: false,
          children: ["fordevelopers/", "fordevelopers/quickconnect", "fordevelopers/clients", "fordevelopers/guides"],
        },
        {
          title: "ERC20",
          collapsable: false,
          children: [
            "modules/erc20/",
            "modules/erc20/01_concepts",
            "modules/erc20/02_state",
            "modules/erc20/03_state_transitions",
            "modules/erc20/04_transactions",
            "modules/erc20/05_hooks",
            "modules/erc20/06_events",
            "modules/erc20/07_parameters",
            "modules/erc20/08_clients",
          ],
        },
        {
          title: "Pools",
          collapsable: false,
          children: [
            "modules/pools/",
            "modules/pools/0x_weights",
            "modules/pools/01_concepts",
            "modules/pools/02_pool_params",
            "modules/pools/03_msgs",
            "modules/pools/04_params",
          ],
        },
        {
          title: "Currency",
          collapsable: false,
          children: [
            "modules/currency/",
          ],
        },
        {
          title: "Infr",
          collapsable: false,
          children: [
            "modules/infr/",
          ],
        },
      ],
    },
    sidebarExtra: {
      "/": [
        {
          title: "Privacy Policy",
          icon: "/icons/icon-book.svg",
          link: "https://imversed.com/privacy.html"
        },
      ]
    },
  },
  markdown: {
    extractHeaders: [ 'h2', 'h3', 'h4', 'h5' ],
    anchor: {
      level: [2, 3, 4]
    },
    toc: {
      format(content) {
        return content.replace(/[_`]/g, "");
      }
    },
    extendMarkdown(md) {
      // provide our own highlight.js to customize Prism setup
      md.options.highlight = require("./theme/highlight");
      // add markdown extensions
      md.use(require("./theme/markup"))
        .use(require("markdown-it-deflist"))
        .use(require("markdown-it-imsize"));
    }
  },
  postcss: {
    plugins: require("../../postcss.config.js").plugins
  }
};
