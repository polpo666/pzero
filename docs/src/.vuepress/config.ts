// @ts-ignore
import { defineUserConfig } from "vuepress";
import theme from "./theme.js";

export default defineUserConfig({
  base: "/pzero/",

  locales: {
    "/": {
      lang: "zh-CN",
      title: "Pzero Framework",
      description: "Pzero docs",
    },
  },

  theme,

  // 和 PWA 一起启用
  // shouldPrefetch: false,
});
