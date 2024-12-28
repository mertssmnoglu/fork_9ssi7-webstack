function transactionList() {
    return {
        transactions: [],
        loading: true,
        error: null,
        async init() {
            try {
                const response = await fetch('/api/transactions')
                this.transactions = await response.json()
            } catch (err) {
                this.error = 'İşlemler yüklenirken bir hata oluştu'
            } finally {
                this.loading = false
            }
        },
        formatAmount(amount) {
            return new Intl.NumberFormat('tr-TR', {
                style: 'currency',
                currency: 'TRY'
            }).format(amount)
        }
    }
}

PetiteVue.createApp().mount("#transaction-list")