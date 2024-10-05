$('#nova-publicacao').on('submit', criarPulicacao)

function criarPulicacao(event) {
    event.preventDefault();

    $.ajax({
        url: "/publicacoes",
        method: "POST",
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val(),
        }
    }).done(function() {
        window.location = "/home"
    }).fail(function() {
        window.location("Deu errado")
    })
}