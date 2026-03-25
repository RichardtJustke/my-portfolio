// Relógio ao vivo (UTC)
const clock = document.getElementById('clock')
if (clock) {
  const tick = () => {
    const now = new Date()
    const h = String(now.getUTCHours()).padStart(2, '0')
    const m = String(now.getUTCMinutes()).padStart(2, '0')
    const s = String(now.getUTCSeconds()).padStart(2, '0')
    clock.textContent = `${h}:${m}:${s} UTC`
  }
  tick()
  setInterval(tick, 1000)
}

// Atalho: pressionar 'e' copia o email
document.addEventListener('keydown', (ev) => {
  // ignora se estiver digitando em algum input
  if (ev.target.tagName === 'INPUT' || ev.target.tagName === 'TEXTAREA') return

  if (ev.key === 'e') {
    const emailLink = document.querySelector('a[href^="mailto:"]')
    if (!emailLink) return
    const email = emailLink.href.replace('mailto:', '')
    navigator.clipboard.writeText(email).then(() => {
      emailLink.textContent = 'copiado!'
      setTimeout(() => { emailLink.textContent = 'email' }, 2000)
    })
  }
})
