import './index.scss'
import useFetch from "../../useFetch";
import {useNavigate, useParams} from "react-router-dom";


function TimeWidget() {
    // d6e2f387-a6ea-471b-96c3-d46a0e7c796d

    // data
    return <div className='time_widget'>
        <div className="time">
            {new Date().toLocaleTimeString('ru-RU', { hour12: false, hour: "numeric", minute: "numeric"})}
        </div>
        <div className="date">
            {new Date().toLocaleString('ru-RU', { weekday: 'short', month: 'long', day: 'numeric' })}
        </div>
    </div>
}


export default TimeWidget
