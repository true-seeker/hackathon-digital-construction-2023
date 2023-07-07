import './index.scss'
import useFetch from "../../useFetch";
import {useNavigate, useParams} from "react-router-dom";
import up from './up.svg'
import down from './down.svg'
import React from "react";

function USDWidget() {
    // 63baeddd-2a07-4f71-aa19-62ecbae26429
    const {data, loading, error} = useFetch(`http://62.84.99.184:80/api/widgets/currency`)
    // data
    return <div className={'usd_widgets'}>
        <div className="content">
            {loading && <>Загрузка</>}
            {data?.currencies?.map((item, index) => {
                return (
                    <div className="row">
                        <div className="col">{item?.name}</div>
                        <div className="col">
                            {!index%2?<img src={up}/>:<img src={down}/>}
                        </div>
                        <div className="col">{item?.value}</div>
                    </div>
                )
            })}
        </div>
    </div>
}


export default USDWidget
