import {createStore} from "vuex"

const store = createStore({
    state() {
        return {
            RSA: '-----BEGIN PUBLIC KEY-----\n' +
                'MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDbzzkYFQeshu39thT4UvpdU7x3\n' +
                'i1FOL6KQ5BJUGPTnVJRS8HU3yuPL5S2tYgPzMJbAm+nPzp7zlpR3Ntop/djET9Sc\n' +
                '2Ne8qCLQyf3ZwsCYxdYbpDVqtWDtlcBNMAaqr2Z4v9Z7ZHfNEWOlkLTvUCXeTT3k\n' +
                '3DFN72wOkP2jLkFJIQIDAQAB\n' +
                '-----END PUBLIC KEY-----\n',
        }
    },
    getters() {
    },
    mutations() {
    },
    actions() {
    },
    modules() {
    },
})

export default store
