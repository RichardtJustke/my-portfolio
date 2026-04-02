// ── Tema ──
const themeBtn = document.getElementById('themeBtn')
const sunIcon = document.getElementById('sunIcon')
const moonIcon = document.getElementById('moonIcon')
const savedTheme = localStorage.getItem('theme') || 'dark'

function updateThemeIcon() {
  const isLight = document.body.classList.contains('light')
  if (isLight) {
    sunIcon.style.display = 'inline'
    moonIcon.style.display = 'none'
  } else {
    sunIcon.style.display = 'none'
    moonIcon.style.display = 'inline'
  }
}

if (savedTheme === 'light') {
  document.body.classList.add('light')
}

updateThemeIcon()

if (themeBtn) {
  themeBtn.addEventListener('click', () => {
    document.body.classList.toggle('light')
    updateThemeIcon()
    const isLight = document.body.classList.contains('light')
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
            const isLight = document.body.classList.contains('light')
            const r = 255
            const g = isLight ? 140 : 107
            const b = isLight ? 66 : 53
            cx.fillStyle = `rgba(${r}, ${g}, ${b}, ${0.15 + d * 0.85})`
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
