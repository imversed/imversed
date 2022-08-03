<template>
  <aside class="left-bar">
      <!-- @apply h-screen overflow-y-auto pb-32;
  height: calc(100vh - 48px); -->
    <div class="wrap h-screen overflow-y-auto pb-16">
      <div id="top" class="mt-8 flex items-center" style="height: 37px;">
        <RouterLink :to="'/'" ref="siteName" class="site-name pl-4 pr-10">
          <img src="/images/logo.svg" />
        </RouterLink>
      </div>

      <slot name="top" />

      <div id="mid" class="relative mt-8">
        <div class="nav-wrap relative w-64 overflow-hidden h-full">
          <div class="set-nav">
            <loading-bars v-if="loading" class="ml-4 mt-1" />
            <transition :name="getTransitionClass(1)">
              <SidebarLinks
                v-if="currentSidebarDepth === 1"
                class="left-bar-links"
                :depth="0"
                :items="items"
                :extra-items="extraItems"
              />
            </transition>
            <transition name="slide-right">
              <SidebarLinks
                v-if="currentSidebarDepth === 2"
                class="left-bar-links"
                :depth="0"
                :items="items"
                :extra-items="extraItems"
              />
            </transition>
          </div>
        </div>
      </div>

      <slot name="bottom" />
    </div>
  </aside>
</template>

<style lang="postcss">
.site-name {
  color: var(--sidebar-link-color);
}

.left-bar-links {

}

.sidebar-transitioning {
  .left-bar {
    transition: all 0.5s cubic-bezier(0.86, 0, 0.07, 1);
  }
}

.left-bar {
  @apply w-64 h-screen fixed z-10;
  background-color: var(--sidebar-bg-color);
  transform: translateX(-16rem);

  .wrap {
    @apply w-64 absolute right-0;
  }

  .left-bar-bottom {
    @apply absolute w-full border-t;
    border-color: var(--border-color);
  }
}

@screen lg {
  .left-bar {
    width: calc(50% - 256px);
  }
}

@screen xl {
  .left-bar {
    width: calc(50% - 384px);
  }
}

.left-bar-bottom {
  @apply h-12;
}

/**
 * transitions
 */

.slide-left-enter-active,
.slide-left-leave-active,
.slide-right-enter-active,
.slide-right-leave-active {
  transition: opacity 350ms cubic-bezier(0.22, 1, 0.36, 1),
    transform 350ms cubic-bezier(0.22, 1, 0.36, 1);
  opacity: 1;
  transform: translateX(0);
}

/* slide from left to target */
.slide-left-enter,
.slide-left-leave-to {
  position: absolute;
  transform: translateX(-16rem);
  opacity: 0;
}

/* slide from right to target */
.slide-right-enter,
.slide-right-leave-to {
  position: absolute;
  transform: translateX(16rem);
  opacity: 0;
}

.slide-up-enter-active {
  transition: opacity 350ms cubic-bezier(0.22, 1, 0.36, 1),
    transform 350ms cubic-bezier(0.22, 1, 0.36, 1);
  opacity: 1;
  transform: translateY(0);
}

/* immediate outtro */
.slide-up-leave-active {
  transition: opacity 0, transform 0;
}

.slide-up-enter,
.slide-up-leave-to {
  transform: translateY(4rem);
  opacity: 0;
}

@media (prefers-reduced-motion: reduce) {
  .sidebar-transitioning .left-bar,
  .slide-left-enter-active,
  .slide-left-leave-active,
  .slide-right-enter-active,
  .slide-right-leave-active,
  .slide-up-enter-active,
  .slide-up-leave-active {
    transition: none;
  }
}
</style>

<script>
import SidebarLinks from "./SidebarLinks.vue";
import LoadingBars from "./LoadingBars.vue";
import {
  resolveSidebarConfig,
  getRelativeActiveBaseFromConfig,
} from "../util";

export default {
  props: ["items", "extraItems", "language"],
  components: { SidebarLinks, LoadingBars },
  data() {
    return {
      currentSidebarDepth: null,
      previousSidebarDepth: null,
      loading: true
    };
  },
  mounted() {
    this.currentSidebarDepth = this.getSidebarNavigationDepth();
    this.loading = false;
  },
  methods: {
    getSidebarNavigationDepth() {
      let config = resolveSidebarConfig(
        this.$page,
        this.$themeConfig
      );

      const sidebarBase = getRelativeActiveBaseFromConfig(this.$page.regularPath, config);

      if (sidebarBase === "/") {
        return 1;
      }

      return 2;
    },
    getTransitionClass(depth) {
      if (!this.previousSidebarDepth) {
        return;
      }

      const direction =
        this.currentSidebarDepth > this.previousSidebarDepth ? "right" : "left";

      // if the depth of this thing is lower and we’re moving right, slide it to the left
      if (direction === "right" && depth < this.currentSidebarDepth) {
        return depth < this.currentSidebarDepth ? "slide-left" : "slide-right";
      }

      // invert if we’re moving to the left
      return depth < this.currentSidebarDepth ? "slide-right" : "slide-left";
    },
  },
  watch: {
    items(items) {
      this.previousSidebarDepth = this.currentSidebarDepth;
      this.currentSidebarDepth = this.getSidebarNavigationDepth();
    },
  },
};
</script>
