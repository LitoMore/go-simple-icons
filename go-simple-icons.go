/**
 * Author: Marco Ollivier
 * File: init.go
 */

package lib

type License struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

type Icon struct {
	Title      string  `json:"title"`
	Slug       string  `json:"slug"`
	Hex        string  `json:"hex"`
	Source     string  `json:"source"`
	Svg        string  `json:"svg"`
	Path       string  `json:"path"`
	Guidelines string  `json:"guidelines"`
	License    License `json:"license"`
}
