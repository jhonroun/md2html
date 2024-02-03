# md2html
Simplest way to convert .md to html with GitHub styilish
```
It used https://api.github.com/markdown
```
Ru:

Простой инструмент для конвертации файлов .md в html с поддержкой стилей GitHub 
```
Использует https://api.github.com/markdown
```
------

The idea arose as a result of the desire to save: https://github.com/Konstantin8105/Effective_Go_RU

**and of course using part of the code:**
* https://github.com/sindresorhus/github-markdown-css
* https://github.com/gilliek/ghmd/tree/master

Ru:

Мысль возникла вследствии желания сохранить: https://github.com/Konstantin8105/Effective_Go_RU

**ну и разумеется с использованием части кода:**
* https://github.com/sindresorhus/github-markdown-css
* https://github.com/gilliek/ghmd/tree/master

------

HowTo (Использование:):

```
git clone https://github.com/jhonroun/md2html
cd md2html
go mod tidy
go build
```

After that:

```
md2html 'example\Эффективный Go.md'
```