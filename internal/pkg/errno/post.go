// Copyright 2024 Yongjun Wang <wangyj641@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/wangyj641/miniblog.

package errno

// ErrPostNotFound 表示未找到博客.
var ErrPostNotFound = &Errno{HTTP: 404, Code: "ResourceNotFound.PostNotFound", Message: "Post was not found."}
