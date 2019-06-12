let g:jekyll_post_extension = '.md'
let g:jekyll_post_template =  [
  \ '---',
  \ 'layout: default',
  \ 'title: "JEKYLL_TITLE"',
  \ 'author: 李佶澳',
  \ 'createdate: "JEKYLL_DATE"',
  \ 'changedate: "JEKYLL_DATE"',
  \ 'categories:',
  \ 'tags:',
  \ 'cover:',
  \ 'keywords:',
  \ 'description:',
  \ '---',
  \ '',
  \ '## 本篇目录',
  \ '',
  \ '* auto-gen TOC:',
  \ '{:toc}',
  \ '',
  \ '## 说明',
  \ '',
  \ '## 参考',
  \ '',
  \ '1. [李佶澳的博客][1]',
  \ '',
  \ '[1]: https://www.lijiaocn.com "李佶澳的博客"',
  \ '']
