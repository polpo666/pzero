// @ts-ignore
import { hopeTheme } from "vuepress-theme-hope";
import { navbarConfig } from "./navbar/index.js";
import { sidebarConfig } from "./sidebar/index.js";

// @ts-ignore
export default hopeTheme({
  hostname: "https://polpo666.github.io/pzero",

  author: {
    name: "polpo666",
    url: "https://github.com/polpo666",
  },

  iconAssets: "iconify",

  copyright: 'Copyright © 2024-2026 polpo666',

  // made by https://gopherize.me
  // favicon.ico made by https://www.bitbug.net
  logo: "/favicon.ico",

  repo: "polpo666/pzero",

  docsDir: "docs/src",

  locales: {
    "/": {
      // 导航栏
      navbar: navbarConfig,

      // 侧边栏
      sidebar: sidebarConfig,

      // 页脚
      footer: "",
      displayFooter: true,

      // Page meta
      metaLocales: {
        editLink: "在 GitHub 上编辑此页",
      },
    },
  },

  // 在这里配置主题提供的插件
  plugins: {
    blog: {
      // category: "category",
      // tag: "tag",
      // star: "star",
    },
    components: {
      components: ["Badge", "VPCard"],
    },

    // 此处开启了很多功能用于演示，你应仅保留用到的功能。
    mdEnhance: {
      align: true,
      attrs: true,
      codetabs: true,
      component: true,
      // demo: true,
      figure: true,
      imgLazyload: true,
      imgSize: true,
      include: true,
      mark: true,
      // stylize: [
      //   {
      //     matcher: "Recommended",
      //     replacer: ({ tag }) => {
      //       if (tag === "em")
      //         return {
      //           tag: "Badge",
      //           attrs: { type: "tip" },
      //           content: "Recommended",
      //         };
      //     },
      //   },
      // ],
      // sub: true,
      // sup: true,
      // tabs: true,
      // tasklist: true,
      // vPre: true,
    },
  },
});
