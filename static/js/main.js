// ── Dot Sphere ──
const cv = document.getElementById('sphere')
if (cv) {
  const cx = cv.getContext('2d')
  let t = 0

  ;(function draw() {
    cx.clearRect(0, 0, 60, 60)
    const R = 24, rows = 12, cols = 12
    for (let i = 0; i < rows; i++) {
      const lat = (i / (rows - 1) - 0.5) * Math.PI
      const y = 30 + R * Math.sin(lat)
      const rr = R * Math.cos(lat)
      for (let j = 0; j < cols; j++) {
        const lon = (j / cols) * Math.PI * 2 + t
        const x = 30 + rr * Math.cos(lon)
        const d = (Math.cos(lon) * Math.cos(lat) + 1) / 2
        if (d > 0.08) {
          cx.beginPath()
          cx.arc(x, y, d * 1.2 + 0.3, 0, Math.PI * 2)
          const color = '59, 130, 246'
          cx.fillStyle = `rgba(${color}, ${0.15 + d * 0.85})`
          cx.fill()
        }
      }
    }
    t += 0.006
    requestAnimationFrame(draw)
  })()
}

// ── Link Previews ──
const previewCard = document.getElementById('previewCard')
const previewLinks = document.querySelectorAll('[data-preview]')

previewLinks.forEach(link => {
  link.addEventListener('mouseenter', () => {
    const imgUrl = link.getAttribute('data-preview')
    previewCard.innerHTML = `<img src="${imgUrl}" alt="Preview">`
    previewCard.classList.add('visible')
  })

  link.addEventListener('mousemove', (e) => {
    const cardW = 280
    const cardH = 180
    let x = e.clientX + 16
    let y = e.clientY + 16

    // Keep card within viewport
    if (x + cardW > window.innerWidth) x = e.clientX - cardW - 16
    if (y + cardH > window.innerHeight) y = e.clientY - cardH - 16

    previewCard.style.left = `${x}px`
    previewCard.style.top = `${y}px`
  })

  link.addEventListener('mouseleave', () => {
    previewCard.classList.remove('visible')
  })
})

// ── Reveal on Scroll ──
const observer = new IntersectionObserver((entries) => {
  entries.forEach((entry, i) => {
    if (entry.isIntersecting) {
      // Stagger the animations
      setTimeout(() => {
        entry.target.classList.add('visible')
      }, i * 60)
    }
  })
}, { threshold: 0.1 })

document.querySelectorAll('.reveal').forEach(el => observer.observe(el))



// ── Stack Easter Eggs ──
document.querySelectorAll('.stack-icon img[data-hover]').forEach(img => {
  img.addEventListener('mouseenter', () => {
    img.src = img.dataset.hover
    img.classList.add('easter-egg-active')
  })
  img.addEventListener('mouseleave', () => {
    img.src = img.dataset.original
    img.classList.remove('easter-egg-active')
  })
})

// ── Temperature (Brasília) ──
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
