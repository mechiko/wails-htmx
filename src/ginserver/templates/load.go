package templates

import (
	"firstwails/domain"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"os"
	"path"
	"strings"
)

// LoadTemplates parses HTML templates and stores them in a cache for efficient rendering during runtime.
func (t *templates) LoadTemplates(debug bool) (err error) {
	t.fs = root
	if debug {
		t.fs = os.DirFS(domain.RootPathTemplateGinDebug)
	}
	embededPages, err := root.ReadDir(".")
	if err != nil {
		return fmt.Errorf("%s %w", modError, err)
	}
	if debug {
		embededPages, err = os.ReadDir(domain.RootPathTemplateGinDebug)
		if err != nil {
			return fmt.Errorf("%s %w", modError, err)
		}
	}
	for i, page := range embededPages {
		t.Logger().Debugf("page %d %s %v", i, page.Name(), page.IsDir())
		if page.IsDir() {
			if err := t.ParsePage(page.Name(), page); err != nil {
				return fmt.Errorf("%s %w", modError, err)
			}
		}
	}
	return nil
}

func (t *templates) ParsePage(page string, dir fs.DirEntry) (err error) {
	// создаем новый шаблон страницы
	if _, ok := t.pages[page]; ok {
		return fmt.Errorf("%s такой шаблон вида %s уже обработан", modError, page)
	}
	t.pages[page] = template.New(page).Funcs(functions)
	embededHtmls, err := root.ReadDir(page)
	if err != nil {
		return fmt.Errorf("%s %w", modError, err)
	}
	for _, html := range embededHtmls {
		// t.Logger().Debugf("htmls %d %s %v", i, html.Name(), html.IsDir())
		if !html.IsDir() {
			if err := t.ParsePageHtml(page, html.Name(), html, t.pages[page]); err != nil {
				return fmt.Errorf("%s %w", modError, err)
			}
		}
	}
	return nil
}

func (t *templates) ParsePageHtml(page string, html string, dir fs.DirEntry, templ *template.Template) error {
	t.Logger().Debugf("parse %s %s %v", page, html, dir.IsDir())
	name, _ := strings.CutSuffix(path.Base(html), path.Ext(html))
	path := path.Join(page, html)
	t.Logger().Debugf("file %s", path)
	if file, err := t.fs.Open(path); err != nil {
		return fmt.Errorf("%s %w", modError, err)
	} else {
		if txt, err := io.ReadAll(file); err != nil {
			return fmt.Errorf("%s %w", modError, err)
		} else {
			t.Logger().Debugf("file read bytes %s %s %d", name, path, len(txt))
		}
	}
	// if txt, err := root.ReadFile(path); err != nil {
	// 	return fmt.Errorf("%s %w", modError, err)
	// } else {
	// 	t.Logger().Debugf("file read bytes  %s %s %d", name, path, len(txt))
	// 	templ.New(name).Funcs(functions).Parse(string(txt))
	// }
	return nil
}
