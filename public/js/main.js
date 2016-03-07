$("#generate").click(function() {

  $.getJSON("/expression", function(data) {
    $("#expression").text(data['phrase']);
    $("#meaning").text(data['meaning']);
  });
});
