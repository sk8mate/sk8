'use strict';

function movedPermanently(value) {
    return {
        status: '301',
        statusDescription: 'Moved Permanently',
        headers: {
            'location': [{
                key: 'Location',
                value,
            }],
            'cache-control': [{
                key: 'Cache-Control',
                value: "max-age=3600"
            }],
        },
    }
}

exports.handler = (event, _, callback) => {
    const request = event.Records[0].cf.request;

    if (request.uri == "/backside") {
        const response = movedPermanently('https://github.com/sk8mate/sk8/tree/main/backside');
        callback(null, response);
    }
    else {
        const response = movedPermanently(`https://github.com/sk8mate/sk8${request.uri}`);
        callback(null, response);
    }
};
