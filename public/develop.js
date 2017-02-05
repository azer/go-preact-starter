remoteControl.register('reload', function () {
  location.reload(true)
})

remoteControl.register('reload-css', function () {
  document.querySelectorAll('link').forEach(style => {
    if (style.rel === "stylesheet") {
      style.href = style.href.replace(/\.css[\?\d]*$/, '.css?' + Date.now())
    }
  })
})
