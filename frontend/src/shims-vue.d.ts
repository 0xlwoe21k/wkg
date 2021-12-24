/* eslint-disable */
declare module '*.vue' {
  import type { DefineComponent } from 'vue'
  '@kangc/v-md-editor/lib/theme/vuepress.js'
  const component: DefineComponent<{}, {}, any>
  export default component
}


declare module '@kangc/v-md-editor/lib/theme/github.js';
declare module 'codemirror';
declare module '@kangc/v-md-editor';
declare module '@kangc/v-md-editor/lib/codemirror-editor';
declare module '@kangc/v-md-editor/lib/preview'