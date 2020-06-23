var all = ""; document.querySelectorAll("[itemprop=\"name codeRepository\"]").forEach(item => all += "git@github.com:Jblew/" + item.innerText + ".git\n"); console.log(all);
