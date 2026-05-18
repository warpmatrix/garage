= Latex

== 小寄巧

压缩论文模板

```latex
% used for usenix
% Compact section/subsection/subsubsection
\usepackage{titlesec}
\titlespacing*{\section}{0pt}{3mm}{2mm}
\titlespacing*{\subsection}{0pt}{2mm}{1mm}
\titlespacing*{\subsubsection}{0pt}{1mm}{0.5mm}

\usepackage{enumitem}
```

小黑圈数字

```latex
\newcommand{\circled}[1]{\tikz[baseline=(myanchor.base)] \node[circle,fill=.,inner sep=0.3pt] (myanchor) {\color{-.}\bfseries\footnotesize #1};}
```
