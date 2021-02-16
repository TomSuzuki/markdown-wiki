// setFormSubmit ...
window.addEventListener("load", function () {
    if (document.getElementById("frame_write") !== null) onsubmit("frame_write", "new");
});

function onsubmit(formId, type) {
    // sendData ...Fromの送信を行う。
    function sendData(form) {
        // data
        let name = document.getElementsByName("name")[0].value

        // object
        const XHR = new XMLHttpRequest();
        const FD = new FormData(form);

        // add event
        XHR.onreadystatechange = () => {
            if (XHR.readyState === 4) {
                if (XHR.status === 200) {
                    location.href = `/page?w=${name}`;
                } else {
                    alert("送信エラーが発生しました。")
                }
            }
        }

        // send
        XHR.open("POST", "/save");
        XHR.send(FD);
        return;
    }

    const form = document.getElementById(formId);
    form.addEventListener("submit", function (event) {
        event.preventDefault();

        // check ...既にページが存在するかチェックする。
        let name = document.getElementsByName("name")[0].value
        accessServer(`/page/status?name=${name}`, (result) => {
            let json = JSON.parse(result);

            // send or error
            if (!json["is_exist"]) sendData(form);
            else alert("既に存在するページです。ページ名を変更してください。");
        });


    });
}

// accessServer ...
function accessServer(path, Callback) {
    let httpObj = new XMLHttpRequest();
    httpObj.onreadystatechange = function () {
        if (httpObj.readyState === 4 && httpObj.status === 200) Callback(httpObj.responseText);
    }
    httpObj.open("GET", path, true);
    httpObj.send(null);
}