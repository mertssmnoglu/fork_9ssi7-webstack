function accountWidget() {
    return {
        account: null,
        async init() {
            const accountEl = document.getElementById('account-data')
            if (!accountEl) {
                return
            }
            this.account = JSON.parse(accountEl.textContent)
        },
        formatAmount(amount) {
            return new Intl.NumberFormat('tr-TR', {
                style: 'currency',
                currency: 'TRY'
            }).format(amount)
        }
    }
}

PetiteVue.createApp().mount("#account-widget")