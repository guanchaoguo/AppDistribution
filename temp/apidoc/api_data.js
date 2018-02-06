define({ "api": [
  {
    "type": "Get",
    "url": "/getCaptcha",
    "title": "获取验证码",
    "description": "<p>获取验证码</p>",
    "group": "captcha",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>token.</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n\tcode: 0,\n\tmessage: \"操作成功\",\n\tresult: - {\n\t\tdata: - {\n\t\t\tcaptcha: \"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAGQAAAAoCAYAAAAIeF9DAAADXklEQVR4nOxbvapUMRDOvcjCdltb2CkoKIi3WBRRtBCsfADBzgcQLYULlooPYCf4AL6AriLIFlcEBQXtLKy3W9hmJYebJY6TZCaZyclZ94PD5vxk5ku+zMw54d59s0NTGKwgd+7eW/fNQQN7fRPIgS/G61cvVcbw8PDJxsezw8fV5mmQgphjUWqI4UAR5eSZs5t+v398z+I22JSlJUYIl67diKZIXwzsnIrBClITs9ms+02JIoEmBTl98WDtjj78w/T06f3bPf9XE83VEEyEn5+PmuOJQaKGNDXQWEQMRZRSkFPWhctX16l2a9DgrD1e0qqjkPjy8UPRCrY+lstl1x6Pxwa2uRGiwdm3WTLe+w8ebey8eP70LzukCHHO7W+onUvOeAO1k28PrM2FBmesr+XuDkNYCL4Y2HmQlDXMIc19PtYfRktNHileUIxUf8gFCmBAlKDEQ6HJvZ6CxuTl8KDY82Ftc2qJzwUTZD6fb54hRwhGAN7nioHZKYWmyBiwOUjNlRXFijCdTjsxfPT2YShVf4ywuJx6AH3BWoUtaPvrRIBiGE7KiqnOmZBaKxj6wFYvtR5gqz43G6Tss4p6bOIpE62ZpkL3uPke6+/7yeWNCYrZI6UsZwwLRQ64aYqaOig2qJB8nYd2Kfb/uZBabdj13JQVa1P950x4Lm/qsyUIpqQQYqJwBpWL1IQ6HqfOnV9PJpPNPYnaJV3/MOz7zqikclNJ6us59jz0H4siK4ZtLxaL7rDn3A+6vtAJAlcapU74opQKRLlH5ePEgHDXc7lK1DMKgq+BkEwqv8f6+3awPqFXytTrN1Z/QoJY/Pr2Fd0CoaJGyjrhGtTVGhPGfxujko8tCMyO9oT0jS5l5e58hlIbFt7cr14qoC8XBRB+dFD4pPxoodqeD0RoBzVkMyYYZsNPXVCkllOWCKAYsOCXFH7K8xo2JftxULy5iE009rqqtbK4QreeskTA/YaRfn6bIiTpmNLOta39fC5f7vglhRLdG6Li4PrNzsdqtUKLe0nRLeUrtTMsCs1VYcWwByz42KHFXcJGtQjRhosOcxwho9HIHL17sycRIduAXv+214rhsO1f41du3SZFUXVBbDTEzrcRTgyqKDtUwE6MgaLJ/w/5n7ETpDH8CQAA//+M1lV7zXAz+QAAAABJRU5ErkJggg==\",\n\t\t\tcaptchaKey: \"wkj4aVGWJKbmXafRQbraWcH0L3da5DZq\"\n\t\t}\n\t}\n}",
          "type": "json"
        }
      ]
    },
    "filename": "app/controllers/captcha.go",
    "groupTitle": "captcha",
    "name": "GetGetcaptcha"
  },
  {
    "type": "Post",
    "url": "/login",
    "title": "后台登录",
    "description": "<p>后台登录</p>",
    "group": "login",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>token.</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locale",
            "description": "<p>语言</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "user_name",
            "description": "<p>用户名</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "password",
            "description": "<p>密码</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "code",
            "description": "<p>验证码</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "captchaKey",
            "description": "<p>验证码key</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"code\": 400,\n  \"message\": \"登录成功\",\n  \"result\": {\n\t\"token\": \"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50Ijoi5byA5Y-R55So5oi3IiwiZGF0ZSI6MTUxMDY0NjYyMSwiaWQiOjEsInVzZXJfbmFtZSI6ImFkbWluIn0.Nd1AuCJyD0CgDPkjp3lljxhyDCBatBMrcCO-lCnE6GE\"\n  }\n}",
          "type": "json"
        }
      ]
    },
    "filename": "app/controllers/login.go",
    "groupTitle": "login",
    "name": "PostLogin"
  },
  {
    "type": "Post",
    "url": "/loginOut",
    "title": "退出登录",
    "description": "<p>退出登录</p>",
    "group": "login",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locale",
            "description": "<p>语言</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "token",
            "description": "<p>token</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n \"code\": 0,\n \"text\": \"操作成功\",\n \"result\": \"\",\n}",
          "type": "json"
        }
      ]
    },
    "filename": "app/controllers/login.go",
    "groupTitle": "login",
    "name": "PostLoginout"
  },
  {
    "type": "Delete",
    "url": "/merge/{id}",
    "title": "删除app合并信息",
    "description": "<p>删除app合并信息</p>",
    "group": "merge",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locale",
            "description": "<p>语言</p>"
          }
        ]
      }
    },
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>token.</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n \"code\": 0,\n \"text\": \"操作成功\",\n \"result\": \"\",\n}",
          "type": "json"
        }
      ]
    },
    "filename": "app/controllers/merge.go",
    "groupTitle": "merge",
    "name": "DeleteMergeId"
  },
  {
    "type": "Get",
    "url": "/merge",
    "title": "app列表获取",
    "description": "<p>app列表管理</p>",
    "group": "merge",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locale",
            "description": "<p>语言</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "app_name",
            "description": "<p>app名称</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "start_date",
            "description": "<p>开始时间</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "end_date",
            "description": "<p>结束时间</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "per_page",
            "description": "<p>每页显示数据条数，默认15条</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "page",
            "description": "<p>当前的所在页码，默认第1页</p>"
          }
        ]
      }
    },
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>token.</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"code\": 0,\n  \"message\": \"操作成功\",\n  \"result\": {\n\t\"current_page\": 1,\n\t\"data\": [\n\t  {\n\t\t\"id\": 1,// 全局唯一标识Id\n\t\t\"title\": \"豌豆荚\",//应用名称\n\t\t\"logo\": \"{host}/uploads/icon/edsfd.png\",// 下载地址\n\t\t\"apk_name\": \"豌豆荚\", // 安卓名称\n\t\t\"apk_id\": \"2\", // 安卓appId\n\t\t\"ipa_name\": \"豌豆荚\", // ios名称\n\t\t\"ipa_id\": \"3\", // iosid\n\t\t\"updated\": \"2017-12-27 10:13\"\n\t  },\n\t],\n\t\"last_page\": 1,\n\t\"per_page\": 10,\n\t\"total\": 7\n  }\n}",
          "type": "json"
        }
      ]
    },
    "filename": "app/controllers/merge.go",
    "groupTitle": "merge",
    "name": "GetMerge"
  },
  {
    "type": "Get",
    "url": "/merge/{id}",
    "title": "获取app合并信息",
    "description": "<p>获取app合并信息</p>",
    "group": "merge",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locale",
            "description": "<p>语言</p>"
          }
        ]
      }
    },
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>token.</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "\t{\n\t  \"code\": 0,\n\t  \"message\": \"操作成功\",\n\t  \"result\": {\n\t\t\"data\": {\n\t\t    \"id\": 1,// 全局唯一标识Id\n          \"type\": \"0\", // 0 安卓 1 ipa\n\t\t\t\"versions\": \"6.0.0\",// 版本号\n\t\t\t\"shotUrl\": \"biet\",// 四位字符串短连接 下载地址 {host}/appDown/biet\n\t\t\t\"downCount\": \"1\",// 已下载次数\n\t\t\t\"allowDown\": \"2\",//允许下载次数\n\t\t\t\"logo\": \"{host}/uploads/icon/edsfd.png\",// logo地址\n\t\t\t\"app_desc\": \"\",// 应用基本介绍\n\t\t}\n\t  }\n\t}",
          "type": "json"
        }
      ]
    },
    "filename": "app/controllers/merge.go",
    "groupTitle": "merge",
    "name": "GetMergeId"
  },
  {
    "type": "Post",
    "url": "/merge",
    "title": "添加合并应用",
    "description": "<p>添加合并应用</p>",
    "group": "merge",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locale",
            "description": "<p>语言</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "title",
            "description": "<p>应用名称</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "apk_name",
            "description": "<p>安卓名称</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "apk_id",
            "description": "<p>安卓appId</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "ipa_name",
            "description": "<p>ios名称</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "ipa_id",
            "description": "<p>ios名称</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "logo",
            "description": "<p>图片地址 通过上传icon接口获取</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "desc",
            "description": "<p>合并后的应用描述</p>"
          }
        ]
      }
    },
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>token.</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n \"code\": 0,\n \"text\": \"操作成功\",\n \"result\": \"\",\n}",
          "type": "json"
        }
      ]
    },
    "filename": "app/controllers/merge.go",
    "groupTitle": "merge",
    "name": "PostMerge"
  },
  {
    "type": "Post",
    "url": "/ploadIco",
    "title": "上传图片",
    "description": "<p>上传图片</p>",
    "group": "merge",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locale",
            "description": "<p>语言</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "file_name",
            "description": "<p>文件名</p>"
          }
        ]
      }
    },
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>token.</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": " {\n  \"code\": 0,\n  \"text\": \"操作成功\",\n\t  \"result\": {\n\t\t\"data\": {\n\t\t  \"img_url\": \"{host}/uopload/icon/edf.png\",\n\t\t}\n }",
          "type": "json"
        }
      ]
    },
    "filename": "app/controllers/merge.go",
    "groupTitle": "merge",
    "name": "PostPloadico"
  },
  {
    "type": "Put",
    "url": "/merge/{id}",
    "title": "修改app合并信息",
    "description": "<p>修改app合并信息</p>",
    "group": "merge",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locale",
            "description": "<p>语言</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "type",
            "description": "<p>0 安卓 1 ipa</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "versions",
            "description": "<p>版本号</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "shotUrl",
            "description": "<p>短网址</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "downCount",
            "description": "<p>已下载次数</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "allowDown",
            "description": "<p>允许下载次数</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "fileUpload",
            "description": "<p>app 上传文件字段</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "app_desc",
            "description": "<p>应用基本介绍</p>"
          }
        ]
      }
    },
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>token.</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n \"code\": 0,\n \"text\": \"操作成功\",\n \"result\": \"\",\n}",
          "type": "json"
        }
      ]
    },
    "filename": "app/controllers/merge.go",
    "groupTitle": "merge",
    "name": "PutMergeId"
  },
  {
    "type": "Post",
    "url": "/release",
    "title": "app发布",
    "description": "<p>app发布</p>",
    "group": "release",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>token.</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locale",
            "description": "<p>语言</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "uploadfile",
            "description": "<p>上传名称</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"code\": 0,\n  \"message\": \"操作成功\",\n  \"result\":\"\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "app/controllers/release.go",
    "groupTitle": "release",
    "name": "PostRelease"
  },
  {
    "type": "Get",
    "url": "/statistics",
    "title": "app统计数据获取",
    "description": "<p>app统计数据获取</p>",
    "group": "statistics",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locale",
            "description": "<p>语言</p>"
          }
        ]
      }
    },
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>token.</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"code\": 0,\n  \"message\": \"操作成功\",\n  \"result\": {\n\t\"data\": [\n\t  {\n\t\t\"scan_count\": \"1\",// 浏览次数\n\t\t\"download_count\": \"\t2\",//下载次数\n\t\t\"upload_count\": \"6\",// 上传次数\n\t\t\"date\": \"2017-12-27 10:13\"// 时间 只获取最近十天统计\n\t  },\n\t],\n  }\n}",
          "type": "json"
        }
      ]
    },
    "filename": "app/controllers/statistics.go",
    "groupTitle": "statistics",
    "name": "GetStatistics"
  },
  {
    "type": "Get",
    "url": "/statistics/down",
    "title": "app下载统计",
    "description": "<p>app下载统计</p>",
    "group": "statistics",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locale",
            "description": "<p>语言</p>"
          }
        ]
      }
    },
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>token.</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"code\": 0,\n  \"message\": \"操作成功\",\n  \"result\": \"\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "app/controllers/statistics.go",
    "groupTitle": "statistics",
    "name": "GetStatisticsDown"
  },
  {
    "type": "Get",
    "url": "/statistics/scan",
    "title": "app浏览统计",
    "description": "<p>app浏览统计</p>",
    "group": "statistics",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locale",
            "description": "<p>语言</p>"
          }
        ]
      }
    },
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>token.</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"code\": 0,\n  \"message\": \"操作成功\",\n  \"result\": \"\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "app/controllers/statistics.go",
    "groupTitle": "statistics",
    "name": "GetStatisticsScan"
  },
  {
    "type": "Delete",
    "url": "/userApp/downStats/{id}",
    "title": "取消下载密码",
    "description": "<p>取消下载密码</p>",
    "group": "userApp",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locale",
            "description": "<p>语言</p>"
          }
        ]
      }
    },
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>token.</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n \"code\": 0,\n \"text\": \"操作成功\",\n \"result\": \"\",\n}",
          "type": "json"
        }
      ]
    },
    "filename": "app/controllers/userApp.go",
    "groupTitle": "userApp",
    "name": "DeleteUserappDownstatsId"
  },
  {
    "type": "Delete",
    "url": "/userApp/{id}",
    "title": "删除app",
    "description": "<p>删除app</p>",
    "group": "userApp",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locale",
            "description": "<p>语言</p>"
          }
        ]
      }
    },
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>token.</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n \"code\": 0,\n \"text\": \"操作成功\",\n \"result\": \"\",\n}",
          "type": "json"
        }
      ]
    },
    "filename": "app/controllers/userApp.go",
    "groupTitle": "userApp",
    "name": "DeleteUserappId"
  },
  {
    "type": "Get",
    "url": "/appDown/{code}",
    "title": "通过短连接下载app",
    "description": "<p>下载app</p>",
    "group": "userApp",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locale",
            "description": "<p>语言</p>"
          }
        ]
      }
    },
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>token.</p>"
          }
        ]
      }
    },
    "filename": "app/controllers/userApp.go",
    "groupTitle": "userApp",
    "name": "GetAppdownCode"
  },
  {
    "type": "Get",
    "url": "/userApp",
    "title": "app列表获取",
    "description": "<p>app列表管理</p>",
    "group": "userApp",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locale",
            "description": "<p>语言</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "type",
            "description": "<p>0 表示 获取全部app类型  1 ：只获取安卓列表 2：只获取ios列表</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "app_name",
            "description": "<p>app名称</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "start_date",
            "description": "<p>开始时间</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "end_date",
            "description": "<p>结束时间</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "per_page",
            "description": "<p>每页显示数据条数，默认15条</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "page",
            "description": "<p>当前的所在页码，默认第1页</p>"
          }
        ]
      }
    },
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>token.</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"code\": 0,\n  \"message\": \"操作成功\",\n  \"result\": {\n\t\"current_page\": 1,\n\t\"data\": [\n\t  {\n\t\t\"id\": 1,// 全局唯一标识Id\n\t\t\"name\": \"豌豆荚\",//应用名称\n\t\t\"logo\": \"{host}/uploads/icon/edsfd.png\",// 下载地址\n\t\t\"type\": \"0\", // 0 安卓 1 ipa\n\t\t\"is_password\": \"是否启动六位下载密码\",\n\t\t\"password\": \"\",// 下载密码\n\t\t\"app_id\": \"NetDragon.Mobile.iPhone.91Space\t\",// 应用ID号码\n\t\t\"versions\": \"6.0.0\",// 版本号\n\t\t\"shotUrl\": \"kdaz\", //短连接\n\t\t\"updated\": \"2017-12-27 10:13\"\n\t  },\n\t],\n\t\"last_page\": 1,\n\t\"per_page\": 10,\n\t\"total\": 7\n  }\n}",
          "type": "json"
        }
      ]
    },
    "filename": "app/controllers/userApp.go",
    "groupTitle": "userApp",
    "name": "GetUserapp"
  },
  {
    "type": "Get",
    "url": "/userApp/{id}",
    "title": "获取账号app信息",
    "description": "<p>获APP信息</p>",
    "group": "userApp",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locale",
            "description": "<p>语言</p>"
          }
        ]
      }
    },
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>token.</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "\t{\n\t  \"code\": 0,\n\t  \"message\": \"操作成功\",\n\t  \"result\": {\n\t\t\"data\": {\n\t\t    \"id\": 1,// 全局唯一标识Id\n          \"type\": \"0\", // 0 安卓 1 ipa\n\t\t\t\"versions\": \"6.0.0\",// 版本号\n\t\t\t\"shotUrl\": \"biet\",// 四位字符串短连接 下载地址 {host}/appDown/biet\n\t\t\t\"downCount\": \"1\",// 已下载次数\n\t\t\t\"allowDown\": \"2\",//允许下载次数\n\t\t\t\"logo\": \"{host}/uploads/icon/edsfd.png\",// logo地址\n\t\t\t\"app_desc\": \"\",// 应用基本介绍\n\t\t}\n\t  }\n\t}",
          "type": "json"
        }
      ]
    },
    "filename": "app/controllers/userApp.go",
    "groupTitle": "userApp",
    "name": "GetUserappId"
  },
  {
    "type": "Put",
    "url": "/userApp/downStats/{id}",
    "title": "app下载密码设置",
    "description": "<p>app下载密码设置</p>",
    "group": "userApp",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locale",
            "description": "<p>语言</p>"
          }
        ]
      }
    },
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>token.</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": " {\n  \"code\": 0,\n  \"text\": \"操作成功\",\n  \"result\": {\n\t\t\"data\": {\n\t\t\t\"password\": \"1234\",// 四位下载密码\n\t\t}\n }",
          "type": "json"
        }
      ]
    },
    "filename": "app/controllers/userApp.go",
    "groupTitle": "userApp",
    "name": "PutUserappDownstatsId"
  },
  {
    "type": "Put",
    "url": "/userApp/{id}",
    "title": "修改保存账号",
    "description": "<p>修改保存账号</p>",
    "group": "userApp",
    "permission": [
      {
        "name": "JWT"
      }
    ],
    "version": "1.0.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locale",
            "description": "<p>语言</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "type",
            "description": "<p>0 安卓 1 ipa</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "versions",
            "description": "<p>版本号</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "shotUrl",
            "description": "<p>短网址</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "downCount",
            "description": "<p>已下载次数</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "allowDown",
            "description": "<p>允许下载次数</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "fileUpload",
            "description": "<p>app 上传文件字段</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "app_desc",
            "description": "<p>应用基本介绍</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n \"code\": 0,\n \"text\": \"操作成功\",\n \"result\": \"\",\n}",
          "type": "json"
        }
      ]
    },
    "filename": "app/controllers/userApp.go",
    "groupTitle": "userApp",
    "name": "PutUserappId"
  }
] });
