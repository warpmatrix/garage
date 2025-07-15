= Typst Quick Start

#set heading(numbering: (..nums) => {
  let sequence = nums.pos()
  let _ = sequence.remove(0)
  if sequence.len() > 0 {
    numbering("1.", ..sequence)
  }
})

== 基础语法

=== 标题格式的配置

- #link("https://typst.app/docs/reference/model/heading/")[官方文档说明]
- 示例：`#set heading(numbering: "1.")`

=== 标注引用

```typ
#figure(supplement: "Algorithm", kind: "algorithm")[  // 标注类型及显示名称
    ...
] <algo>  // 标注
@algo  // 引用
```

=== 引入第三方库：`#import`

通过 `<lib>.<module>` 调用库的内容。示例：
```typ
#import "@preview/note-me:0.3.0"
#note-me.note[
    ...
]
```

import 所有内容直接使用：
```typ
#import "xx/xx:xx": *
```

== Typst 常用起手式

章节标号（跳过一级标题，#link("https://github.com/typst/typst/discussions/3341#discussioncomment-8364172")[reference]）
```typ
#set heading(numbering: (..nums) => {
  let sequence = nums.pos()
  let _ = sequence.remove(0)
  if sequence.len() > 0 {
    numbering("1.", ..sequence)
  }
})
```

=== 常用的第三方库

checklist 的渲染:
```typ
#import "@preview/cheq:0.2.2": checklist
#show: checklist
```

admonitions 的渲染：
```typ
#import "@preview/gentle-clues:1.2.0"
#import "@preview/note-me:0.4.0" // GitHub-style Admonitions
```

pintora 的渲染：
```typ
#import "@preview/pintorita:0.1.4"
#show raw.where(lang: "pintora"): it => pintorita.render(it.text)
```

伪代码的渲染：
```typ
#import "@preview/algo:0.3.6"
#algo.algo()[
    ...
]
```

代码美化与标注：
````typ
#import "@preview/codly:1.3.0"
#show: codly.codly-init.with()
#codly.codly()
```py
print("hello")
```
````

文字与图片混合布局：
```typ
#import "@preview/wrap-it:0.1.1"
```
