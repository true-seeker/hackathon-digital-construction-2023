import {useEffect, useState} from "react"


function useFetch(url) {
    const [data, setData] = useState(null)
    const [error, setError] = useState(null)
    const [loading, setLoading] = useState(false)

    useEffect(() => {
        setLoading(true);
        fetch(url, {
            method: 'get',
            headers: {
                'Accept': 'application/json, text/plain, */*',
                'Content-Type': 'application/json',
            },
        })
            // .then(response => response.json())
            .then(response => {
                const isJson = response.headers.get('content-type')?.includes('application/json');
                const data = isJson ? response.json() : null;

                if (!response.ok) {
                    data.then(x=>{
                        const error = (x && x.message) || (x && x.error) || response.status;
                        return Promise.reject(error);
                    })
                    return null
                }
                return data
            })

            .then(setData)
            .catch(setError)
            .finally(() => setLoading(false));
    }, [url])

    return {data, error, loading}
}
export default useFetch



// const {data, loading, error} = useFetch('http://62.84.99.184:8080/api/get')
//
// if (error) {
//     console.log(error)
//     return <div>Error</div>
// }
// if (loading) {
//     return <div>Loading...</div>
// }
//
// if (data) {
//      return <>{data.json_value}</>
// }
