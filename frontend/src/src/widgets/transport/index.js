import './index.scss'
import useFetch from "../../useFetch";
import {useNavigate, useParams} from "react-router-dom";
import bus_icon from './bus-icon.svg'
import arrow from './arrow.svg'

function TransportWidget() {
    // e6b16a02-3d14-4185-b02b-ef1c3035f159
    return (
        <div className='transport_widgets'>
            <div className="transport_card">
                <div className="transport_card-header">Прибытие транспорта</div>
                <div className="transport_card-line"></div>
                <div className="transport_card-body">
                    <div className="headers">
                        <span>Маршрут</span>
                        <span>Направление</span>
                        <span>Через</span>
                    </div>

                    <div className="transport-group">
                        <div className="transport-group-header">Автобусы</div>
                        <div className="transport_line">
                            <div className="number">545</div>
                            <div className="icons">
                                <img src={bus_icon}/>
                                <img src={arrow}/>
                            </div>
                            <span>ост. Следующая</span>
                            <span>3 мин</span>
                        </div>
                        <div className="transport_line">
                            <div className="number">545</div>
                            <div className="icons">
                                <img src={bus_icon}/>
                                <img src={arrow}/>
                            </div>
                            <span>ост. Пермская</span>
                            <span>2.1 мин</span>
                        </div>
                        <div className="transport_line">
                            <div className="number">545</div>
                            <div className="icons">
                                <img src={bus_icon}/>
                                <img src={arrow}/>
                            </div>
                            <span>ост. Усадебная</span>
                            <span>3 мин</span>
                        </div>
                    </div>
                    <div className="transport-group">
                        <div className="transport-group-header">Трамваи</div>
                        <div className="transport_line">
                            <div className="number">545</div>
                            <div className="icons">
                                <img src={bus_icon}/>
                                <img src={arrow}/>
                            </div>
                            <span>ост. Ленина</span>
                            <span>1.5 мин</span>
                        </div>
                    </div>
                    <div className="transport-group">
                        <div className="transport-group-header">Метро</div>
                        <div className="transport_line">
                            <div className="number">545</div>
                            <div className="icons">
                                <img src={bus_icon}/>
                                <img src={arrow}/>
                            </div>
                            <span>ост. Пушкина</span>
                            <span>3 мин</span>
                        </div>
                    </div>

                </div>
            </div>
        </div>
    )
}


export default TransportWidget
