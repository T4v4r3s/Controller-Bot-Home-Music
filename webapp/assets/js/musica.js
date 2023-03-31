function criarMusica(evento){
    evento.preventDefault();

    $.ajax({
        url: "/usuarios",
        method: "POST",
        data: $('#n').val(),
    }).done(function() { //201 200 204
        alert("Usuário cadastrado com sucesso!");
    }).fail(function(erro){ //400 404 401 403 404
        alert(erro);
    });
}