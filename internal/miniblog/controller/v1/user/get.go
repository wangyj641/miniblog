// Copyright 2024 Yongjun Wang <wangyj641@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/wangyj641/miniblog.

package user

import (
	"github.com/gin-gonic/gin"

	"github.com/wangyj641/miniblog/internal/pkg/core"
	"github.com/wangyj641/miniblog/internal/pkg/log"
)

// Get 获取一个用户的详细信息.
func (ctrl *UserController) Get(c *gin.Context) {
	log.C(c).Infow("Get user function called")

	user, err := ctrl.b.Users().Get(c, c.Param("name"))
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, user)
}
