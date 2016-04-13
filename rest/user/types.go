/*
Copyright (C) 2016 Paolo Galeone <nessuno@nerdz.eu>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package user

import (
	"github.com/nerdzeu/nerdz-api/nerdz"
)

// Informations represents the user information
type Informations struct {
	Info     *nerdz.InfoTO         `json:"info"`
	Contacts *nerdz.ContactInfoTO  `json:"contacts"`
	Personal *nerdz.PersonalInfoTO `json:"personal"`
}