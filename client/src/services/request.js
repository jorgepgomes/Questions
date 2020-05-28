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

    get = (url) => {
        return new Promise(async (resolve, reject) => {
            axios
                .get(this.makeUrl(url)).then((resp) => {
                    resolve(resp);
                })
                .catch((err) => {
                    reject(err);
                });
        });
    };

    post = (url, data) => {
        return new Promise(async (resolve, reject) => {
            axios
                .post(this.makeUrl(url), data).then((resp) => {
                    resolve(resp);
                })
                .catch((err) => {
                    reject(err);
                });
        });
    };
}

export default Request;