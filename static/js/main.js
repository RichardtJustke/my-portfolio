// ── Tema ──
const themeBtn = document.getElementById('themeBtn')
const savedTheme = localStorage.getItem('theme') || 'dark'

if (savedTheme === 'light') {
  document.body.classList.add('light')
}

if (themeBtn) {
  themeBtn.addEventListener('click', () => {
    const isLight = document.body.classList.toggle('light')
    localStorage.setItem('theme', isLight ? 'light' : 'dark')
  })
}

// ── Esfera animada ──
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

// ── Temperatura via wttr.in (Brasília) ──
const tempEl = document.getElementById('temp')
if (tempEl) {
  fetch('https://wttr.in/Brasilia?format=%t')
    .then(r => r.text())
    .then(d => { 
      tempEl.innerHTML = `<i class="fas fa-cloud"></i> ${d.trim()}`
    })
    .catch(() => { 
      tempEl.innerHTML = `<i class="fas fa-cloud"></i> —`
    })
}
