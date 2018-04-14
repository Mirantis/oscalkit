// oscalkit - OSCAL conversion utility
// Written in 2017 by Andrew Weiss <andrew.weiss@docker.com>

// To the extent possible under law, the author(s) have dedicated all copyright
// and related and neighboring rights to this software to the public domain worldwide.
// This software is distributed without any warranty.

// You should have received a copy of the CC0 Public Domain Dedication along with this software.
// If not, see <http://creativecommons.org/publicdomain/zero/1.0/>.

package oscal

import (
	"encoding/json"
)

// Param ...
type Param struct {
	ID            string `xml:"id,attr,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	OptionalClass string `xml:"class,attr,omitempty" json:"class,omitempty" yaml:"class,omitempty"`
	Desc          Desc   `xml:"desc" json:"desc" yaml:"desc"`
	Label         string `xml:"label,omitempty" json:"label,omitempty" yaml:"label,omitempty"`
	Value         string `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
}

// Desc ...
type Desc struct {
	Raw string `xml:",innerxml"`
}

// MarshalJSON ...
func (d *Desc) MarshalJSON() ([]byte, error) {
	return json.Marshal(formatRawProse(d.Raw))
}

// MarshalYAML ...
func (d *Desc) MarshalYAML() (interface{}, error) {
	return d.Raw, nil
}

// UnmarshalJSON ...
func (d *Desc) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	d.Raw = raw

	return nil
}