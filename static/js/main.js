// Esfera animada
const cv = document.getElementById('sphere')
if (cv) {
  const cx = cv.getContext('2d')
  let t = 0
    ; (function draw() {
      cx.clearRect(0, 0, 60, 60)
      const R = 27, rows = 12, cols = 12
      for (let i = 0; i < rows; i++) {
        const lat = (i / (rows - 1) - 0.5) * Math.PI
        const y = 30 + R * Math.sin(lat)
        const rr = R * Math.cos(lat)
        for (let j = 0; j < cols; j++) {
          const lon = (j / cols) * Math.PI * 2 + t
          const x = 30 + rr * Math.cos(lon)
          const d = (Math.cos(lon) * Math.cos(lat) + 1) / 2
          if (d > 0.05) {
            cx.beginPath()
            cx.arc(x, y, d * 1.4 + 0.3, 0, Math.PI * 2)
            cx.fillStyle = `rgba(61,111,168,${0.15 + d * 0.85})`
            cx.fill()
          }
        }
      }
      t += 0.007
      requestAnimationFrame(draw)
    })()
}

// Temperatura via wttr.in
const tempEl = document.getElementById('temp')
if (tempEl) {
  fetch('https://wttr.in/Sao+Paulo?format=%t')
    .then(r => r.text())
    .then(d => { tempEl.textContent = d.trim() })
    .catch(() => { tempEl.textContent = '' })
}

// Toggle de tema (só visual por enquanto)
const themeBtn = document.getElementById('themeBtn')
if (themeBtn) {
  themeBtn.addEventListener('click', () => {
    document.body.classList.toggle('light')
  })
}
