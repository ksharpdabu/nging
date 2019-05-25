/*
   Nging is a toolbox for webmasters
   Copyright (C) 2018-present  Wenhui Shen <swh@admpub.com>

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published
   by the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package navigate

//Item 操作
type Item struct {
	DisplayOnMenu bool
	Name          string
	Action        string
	Icon          string
	Target        string
	Children      List
}

//List 操作列表
type List []*Item

func (a *List) Remove(index int) List {
	if index < 0 {
		*a = (*a)[0:0]
		return *a
	}
	size := len(*a)
	if size > index {
		if size > index+1 {
			*a = append((*a)[0:index], (*a)[index+1:]...)
		} else {
			*a = (*a)[0:index]
		}
	}
	return *a
}

func (a *List) Set(index int, list ...*Item) List {
	if len(list) == 0 {
		return *a
	}
	if index < 0 {
		*a = append(*a, list...)
		return *a
	}
	size := len(*a)
	if size > index {
		(*a)[index] = list[0]
		if len(list) > 1 {
			a.Set(index+1, list[1:]...)
		}
		return *a
	}
	for start := size; start < index; start++ {
		*a = append(*a, nil)
	}
	*a = append(*a, list...)
	return *a
}

//Add 添加列表项
func (a *List) Add(index int, list ...*Item) List {
	if len(list) == 0 {
		return *a
	}
	if index < 0 {
		*a = append(*a, list...)
		return *a
	}
	size := len(*a)
	if size > index {
		list = append(list, (*a)[index])
		(*a)[index] = list[0]
		if len(list) > 1 {
			a.Add(index+1, list[1:]...)
		}
		return *a
	}
	for start := size; start < index; start++ {
		*a = append(*a, nil)
	}
	*a = append(*a, list...)
	return *a
}
