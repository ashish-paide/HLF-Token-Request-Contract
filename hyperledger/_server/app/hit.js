const http = require('http'); 


const hiturl = (url, method, payload) => {
  console.log("gonna hit url", url, "method", method);
  return new Promise((resolve, reject) => {
    const options = {
      method: method,
      headers: {
        'Content-Type': 'application/json',
      },
    };

    const request = http.request(url, options, (response) => {
      let responseData = '';

      response.on('data', (chunk) => {
        responseData += chunk;
      });

      response.on('end', () => {
        try {
          const parsedResponse = JSON.parse(responseData);
          console.log(responseData);
          resolve(parsedResponse); // Resolve the Promise with the parsed response
        } catch (error) {
          console.error('Error occurred:', error);
          reject(error); // Reject the Promise with the error
        }
      });
    });

    request.on('error', (error) => {
      console.error('Error occurred:', error);
      reject(error); // Reject the Promise with the error
    });

    if (payload) {
      request.write(JSON.stringify(payload));
    }
    request.end();
  });
};

exports.hiturl = hiturl;
