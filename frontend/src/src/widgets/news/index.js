import './index.scss'
import useFetch from "../../useFetch";
import {useNavigate, useParams} from "react-router-dom";
import {useEffect} from "react";


function NewsWidget() {
    // b71bef49-574e-4354-867c-ca77794172be
    const params = useParams()

    const {data, loading, error} = useFetch(`http://62.84.99.184:80/api/widgets/news/${params.screen_id}`)

    if (loading) {
        return <>
            Is loading
        </>
    }
    if (error) {
        return <>
            Some errors
        </>
    }

    // data
    return <div className='news_widgets'>
        <div className="card">
            <div className="news_card_header">Новости Вашего дома</div>

            <div className="line"></div>

            <div className="body">
                {data?.news?.map(item => {
                    return (
                        <div className="message">
                            <div className="news-header">
                                <div className="time">07.07</div>
                                <div className="text">{item?.title}</div>
                            </div>
                            <div className="content" dangerouslySetInnerHTML={{__html: item?.text}}>
                            </div>
                        </div>
                    )
                })}

            </div>
        </div>
    </div>
}


export default NewsWidget
