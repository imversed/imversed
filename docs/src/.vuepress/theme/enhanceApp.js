import Vuex from "vuex";
import CodeToggle from "./components/CodeToggle";
import PreHeading from "./components/PreHeading";
import PostHeading from "./components/PostHeading";

import { setStorage } from "./Storage";

export default ({ Vue, options, siteData }) => {
  const base = siteData.base;

  Vue.component("code-toggle", CodeToggle);
  Vue.component("pre-heading", PreHeading);
  Vue.component("post-heading", PostHeading);

  Vue.use(Vuex);

  Vue.mixin({
    data() {
      return {
        version: null
      };
    },
    computed: {
      $title() {
        const page = this.$page;

        // completely override title from frontmatter
        if (page.frontmatter.title) {
          return page.frontmatter.title;
        }

        // get explicit (frontmatter) or inferred page title
        const pageTitle = page.title ? page.title.replace(/[_`]/g, "") : "";

        // doc set title, global site title, or fall back to `VuePress`
        const siteTitle = (
          this.$siteTitle ||
          "VuePress"
        );

        if (pageTitle && siteTitle && pageTitle !== siteTitle) {
          return `${pageTitle} | ${siteTitle}`;
        }

        return siteTitle;
      },
      $siteTitle() {
        return this.$themeConfig.title || this.$site.title || "";
      },
    },
  });

  Object.assign(options, {
    data: {
      codeLanguage: null
    },

    store: new Vuex.Store({
      state: {
        codeLanguage: null
      },
      mutations: {
        changeCodeLanguage(state, language) {
          state.codeLanguage = language;
          setStorage("codeLanguage", language, base);
        }
      }
    })
  });
};
