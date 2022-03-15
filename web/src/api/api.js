import axios from 'axios'

class Api {
    constructor(p) {
        var baseurl = ""
        this.p = p
        this.axios = axios.create({
            baseURL: baseurl,
            withCredentials: true,
            timeout: 60000
        });
        this.hasLogin = false
    }
    getUserList(params) {
        let t = this
        return new Promise((resolve, reject) => {
            t.axios.get('/api/admin/users', {params: params}).then(resp => {
                resolve(resp)
            }).catch(error => {
                reject(error)
            })
        })
    }
    updateUser(data) {
        let t = this
        return new Promise((resolve, reject) => {
            t.axios.post('/api/admin/user_info', data).then(resp => {
                resolve(resp)
            }).catch(error => {
                reject(error)
            })
        })
    }
    deleteUser(id) {
        let t = this
        return new Promise((resolve, reject) => {
            t.axios.delete('/api/admin/user_info/' + id).then(resp => {
                resolve(resp)
            }).catch(error => {
                reject(error)
            })
        })
    }
}

export default Api