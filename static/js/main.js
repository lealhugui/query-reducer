function query() {
    $.when($.ajax({
      url: "/query",
      type: "POST",
      dataType: "json",
      data: JSON.stringify({queryText: codeEditor.getValue()}),
      contentType: "application/json"
    })).then(function (data) {
      $("#records").html(`<tr></tr>`);
      
      const firstRec = data[0];
      const columns = Object.keys(firstRec);
      let line = "";
      for (k of columns) {
        line += `<td>${k.toUpperCase()}</td>`
      }
      $("#records tr:last").after(`<tr class="header">${line}</tr>`);
      
      for (r of data) {
        line = ""
        for (k of columns) {
          line += `<td>${r[k]}</td>`
        }
        $("#records tr:last").after(`<tr>${line}</tr>`);
      }
      
    }).catch(function(err) {
        console.error(err);
        alert(err);
    });
}