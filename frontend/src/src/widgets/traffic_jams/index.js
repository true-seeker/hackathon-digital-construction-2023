import './index.scss'

function TrafficJamsWidget() {
    return (
        <div className='trafficjams_widgets'>
            <div className="trafficjams_card">
                <div className="trafficjams_card-header">Пробки <div className="trafficjams--red">4</div> баллов</div>
                <div className="trafficjams_card-line"></div>
                <div className="trafficjams_card-body">

                    <div className="trafficjams-group">
                        <div className="trafficjams-group-header">Северный проезд</div>
                        <div className="trafficjams_line">
                            <div className="title">Улица Пушкина</div>
                            <div className="number">5</div>
                        </div>
                        <div className="trafficjams_line">
                            <div className="title">Улица Ленина</div>
                            <div className="number">6</div>
                        </div>
                    </div>

                    <div className="trafficjams-group">
                        <div className="trafficjams-group-header">Южный проезд</div>
                        <div className="trafficjams_line">
                            <div className="title">Ул Усадебная</div>
                            <div className="number number-green">1</div>
                        </div>
                        <div className="trafficjams_line">
                            <div className="title">Ул Пермская</div>
                            <div className="number number-green">2</div>
                        </div>
                    </div>

                    <div className="trafficjams-group">
                        <div className="trafficjams-group-header">Западный проезд</div>
                        <div className="trafficjams_line">
                            <div className="title">Ул Куйбышева</div>
                            <div className="number number-green">4</div>
                        </div>
                        <div className="trafficjams_line">
                            <div className="title">Ул Карпинского</div>
                            <div className="number">6</div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}


export default TrafficJamsWidget
