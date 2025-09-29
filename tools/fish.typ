= Fish

#show link: it => {
  set text(blue)
  underline(it)
}

#link("https://github.com/fish-shell/fish-shell/issues/540")[Fish doesn't support heredoc], but there is other workaround, for example:

```fish
printf "\
...
" | cat
```
