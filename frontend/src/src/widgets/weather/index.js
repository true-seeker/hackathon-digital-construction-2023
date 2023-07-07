import './index.scss'
import useFetch from "../../useFetch";
import {useNavigate, useParams} from "react-router-dom";
import cloud from './cloud.svg'
import sun from './sun.svg'


function WeatherWidget() {
    // 63baeddd-2a07-4f71-aa19-62ecbae26429
    const params = useParams()

    const {
        data,
        loading,
        error
    } = useFetch(`http://62.84.99.184:80/api/widgets/weather/${params.screen_id}`)
    if (loading) {
        return <>Загрузка</>
    }

    // data
    return <div className='weather_widgets'>
        <div className="weather">
            <h2>Сейчас</h2>
            <div className="info">
                <div className="now">
                    <div className="icon">
                        {data?.weather?.condition === 'clear' ? <img src={sun}/> : <img src={cloud}/>}
                    </div>
                    <span>{data?.weather?.temperature_now?.value}°</span>
                </div>
                <div className="bottom">
                    <div className="pressure">
                        <span>{data?.weather?.pressure}</span>
                        <span className='text'>мм<br/>рт ст</span>
                    </div>
                    <div className="cloudy">{data?.weather?.condition === 'clear' ? "Ясно" : "Пасмурно"}</div>
                    <div className="feels">
                        <span>Ощущается как</span>
                        <div className="cels">{data?.weather?.feels_like}°</div>
                    </div>
                </div>
            </div>
        </div>

        <div className="line"></div>
        <div className="next">

            {data?.weather?.Forecast?.map(item => {
                return (
                    <div className="block">
                        <div className="date">{item?.date?.split('-')[2]}</div>
                        <div className="icon">
                            {item?.condition === 'clear' ? <img src={sun}/> : <img src={cloud}/>}
                        </div>
                        <div className="value">{item?.value}°</div>
                    </div>
                )
            })}
        </div>
    </div>
}


export default WeatherWidget
