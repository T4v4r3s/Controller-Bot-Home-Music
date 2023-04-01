$('#formulario-musica').on('submit', criarMusica);

function criarMusica(evento){
    evento.preventDefault();

    console.log($('#url').val())

    $.ajax({
        url: "/addmusica",
        method: "POST",
        data: {
            URL: $('#url').val()
        }
    }).done(function() { //201 200 204
        alert("Música adicionada na fila com sucesso!");
    }).fail(function(erro){ //400 404 401 403 404
        alert(erro);
    });
}