import axios from "axios";

class Request {
    constructor(base_url) {
        this.setBaseUrl(base_url);
    }

    setBaseUrl = (base_url) => {
        this.base_url =
            base_url[base_url.length - 1] === "/"
                ? base_url.substring(0, base_url.length - 2)
                : base_url;

    };

    getBaseUrl = () => {
        return this.base_url;
    };

    makeUrl = (url) => {
        if (!url) return this.getBaseUrl();

        url = url[0] === "/" ? url : "/" + url;
        return this.getBaseUrl() + url;
    };

    makeData = (data) => {
        return data || {};
    };

    request = (method = "get", url, data) => {
        return new Promise(async (resolve, reject) => {
            axios
                .request({
                    method,
                    url: this.makeUrl(url),
                    data: this.makeData(data),
                })
                .then((resp) => {
                    resolve(resp);
                })
                .catch((err) => {
                    reject(err);
                });
        });
    };

    get = (url, data) => {
        return this.request("get", url, data);
    };

    post = (url, data) => {
        return this.request("post", url, data);
    };
}

export default Request;
