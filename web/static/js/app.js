document.addEventListener('DOMContentLoaded', function () {
    // Inputs del formulario
    const textInput = document.getElementById('text');    
    const morseInput = document.getElementById('morse');  
    
    // Variable para manejar retrasos y evitar enviar muchas peticiones seguidas
    let timeout; 

    // Función principal para traducir contenido
    function traducir(inputType, valor) {
        // Cancela cualquier timeout previo para evitar peticiones innecesarias
        clearTimeout(timeout);

        // Establecer un retraso para no enviar una petición en cada tecla presionada
        timeout = setTimeout(() => {
            // Prepara los datos según si se está escribiendo texto o Morse
            const data = inputType === 'text'
                ? { texto: valor }      
                : { morse: valor };      

            // Enviar una petición POST al backend para traducir
            fetch('/traducir', {
                method: 'POST',
                // Indicamos que enviamos JSON
                headers: { 'Content-Type': 'application/json' },
                // Convertimos los datos a JSON 
                body: JSON.stringify(data) 
            })
            // Esperamos la respuesta y la convertimos en JSON
            .then(res => res.json()) 
            .then(res => {
                // Actualiza el otro campo con el resultado recibido del servidor
                if (inputType === 'text') {
                    morseInput.value = res.resultado;
                } else {
                    textInput.value = res.resultado;
                }
            });
        }, 200); // Esperamos 200ms antes de enviar la petición (mejora el rendimiento)
    }

    // Detecta cuando el usuario escribe en el campo de texto
    textInput.addEventListener('input', () => {
        traducir('text', textInput.value); // Llama a traducir con el valor del texto
    });

    // Detecta cuando el usuario escribe en el campo de Morse
    morseInput.addEventListener('input', () => {
        // Llama a traducir con el valor del Morse
        traducir('morse', morseInput.value); 
    });
});
