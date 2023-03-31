$('#formulario-cadastro').on('submit', criarUsuario); //quando um formulario com esse nome receber um submit ele faz essa função
// $ -> pega dados

function criarUsuario(evento){
    evento.preventDefault();

    if($('#senha').val() != $('#confirmar-senha').val()){ //para pegar valor do elemento $('referencia ao elemento').val
        alert("As senhas não coincidem!");
        return
    }

    $.ajax({
        url: "/usuarios",
        method: "POST",
        data:{
            nome: $('#nome').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
            senha: $('#senha').val()
        }
    }).done(function() { //201 200 204
        alert("Usuário cadastrado com sucesso!");
    }).fail(function(erro){ //400 404 401 403 404
        alert(erro);
    });
}