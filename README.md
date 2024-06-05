https://go.dev/doc/tutorial/web-service-gin
https://go.dev/doc/articles/wiki/

GO官方教程精彩卓绝的WEB开发指导

我的精力主要放在了*gowiki*的钻研上

并且我完成了课后作业

```
Other tasks
Here are some simple tasks you might want to tackle on your own:

1. Store templates in tmpl/ and page data in data/.
2. Add a handler to make the web root redirect to /view/FrontPage.
3. Spruce up the page templates by making them valid HTML and adding some CSS rules.
4. Implement inter-page linking by converting instances of [PageName] to
<a href="/view/PageName">PageName</a>. (hint: you could use regexp.ReplaceAllFunc to do this)
```

前3个任务我自信能够理解官方的意图，但是第4个任务确实让我感到费解

将pagename转化为a连接，我的第一感觉是做那种可以实现页面跳转的东西，比如上一页，下一篇文章之类的，但是他们给出的提示却是一个正则替换的函数，我需要替换什么？

然后我想难道是将标题变成可以点击的超链接，考虑了一下后发现没有必要。

这个问题的描述确实让我费解，可能是我的阅读理解还不够好

最后我决定做个目录的跳转吧