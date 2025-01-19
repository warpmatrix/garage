= Typst Quick Start

== 引入第三方库：`#import`

- 通过 `<lib>.<module>` 调用库的内容。示例：
  ```typ
  #import "@preview/note-me:0.3.0"
  #note-me.note[]
  ```
- `#import "xx/xx:xx": *` import 所有模块，可以直接使用

== 标题格式的配置

- #link("https://typst.app/docs/reference/model/heading/")[官方文档说明]
- 示例：`#set heading(numbering: "1.")`
- #link("https://github.com/typst/typst/discussions/3341#discussioncomment-8364172")[更高级的用法]
