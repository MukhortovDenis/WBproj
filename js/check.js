window.addEventListener('DOMContentLoaded', () => {
  
    let resultJSON;
  
    const sendData = async (url, data) => {
      const response = await fetch(url, {
        method: 'POST',
        body: data,
        headers: {
          'Content-Type': 'application/json;charset=utf-8'
          },
      });
  
      resultJSON = await response.json();
  
      return resultJSON;
    };
  
    const sendForm = () => {
      const form = document.getElementById('check');
  
      form.addEventListener('submit', e => {
        e.preventDefault();
  
        const formData = new FormData(form);
        const data = {};
  
        for (const [key, value] of formData) {
          data[key] = value;
        }
        sendData('/check-order', JSON.stringify(data))
            .then(() => {
                  window.location.href = '/result';
              }
            )
            .catch((err) => {
              console.log(err);
            });
      });
    };

    sendForm();
  });