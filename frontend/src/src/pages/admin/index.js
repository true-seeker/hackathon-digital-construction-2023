import './index.scss'
import useFetch from "../../useFetch";
import React, {Component, useEffect, useState} from "react";
import chevron_down from './chevron-down.svg'
import chevron_up from './chevron-up.svg'
import dot from './dot.svg'
import Editor from "./editor";
import {useNavigate, useParams} from "react-router-dom";
import play from './play.svg'
import logo from './1.jpg'
// function Create({setRandom}) {
//     const [name, setName] = useState('')
//     const [address, setAddress] = useState('')
//
//     function handleSubmit(event) {
//         fetch(`http://62.84.99.184:80/api/zhks`, {
//             method: 'post',
//             headers: {
//                 'Accept': 'application/json, text/plain, */*',
//                 'Content-Type': 'application/json',
//             },
//             body: JSON.stringify({
//                 'name': name,
//                 'address': address,
//             })
//         })
//             // .then(response => response.json())
//             .then(response => {
//                 console.log(response)
//             })
//     }
//
//
//     return (
//         <>
//             <tr>
//                 <td>
//                     <input type={"text"} required value={name} onChange={event => setName(event.target.value)}
//                            placeholder={'Name'}/>
//                 </td>
//                 <td>
//                     <button type={'submit'} onClick={handleSubmit}>Save</button>
//                 </td>
//             </tr>
//         </>
//     )
// }

function Screen({id, name}) {
    const navigate = useNavigate();

    function handler(event) {
        event.stopPropagation()
        navigate(`/admin/${id}`);
    }

    return <div onClick={handler}>
        <div className={'item'}>
            <img src={dot} style={{width: '5px'}}/>
            <span>{name}</span>
        </div>
    </div>
}


function Screens({elevator_id}) {
    const {data, loading, error} = useFetch(`http://62.84.99.184:80/api/elevators/${elevator_id}/screens`)
    if (loading) {
        return <>Загрузка</>
    }

    return (
        <ul>
            {data?.screens ? data?.screens?.map(item => {
                return <li><Screen id={item?.id} name={item?.name}/></li>
            }) : <li className={'no'}>Пусто</li>}
        </ul>
    )
}

function Elevator({id, name, address, building_id}) {
    const [collapsed, setCollapsed] = useState(false)

    function handler(event) {
        event.stopPropagation()
        setCollapsed(!collapsed)
    }

    return <div className={'zhk'} onClick={handler}>
        <div className={'item'}>
            {collapsed ? <img src={chevron_down}/> : <img src={chevron_up}/>}
            <span>{name}</span>
        </div>
        {collapsed ? <Screens elevator_id={id}/> : <></>}
        {/*<Buildings zhk_id={id}/>*/}
    </div>
}


function Elevators({building_id}) {
    const {data, loading, error} = useFetch(`http://62.84.99.184:80/api/buildings/${building_id}/elevators`)
    if (loading) {
        return <>Загрузка</>
    }

    return (
        <ul>
            {data?.elevators ? data?.elevators?.map(item => {
                return <li><Elevator id={item?.id} name={item?.name} address={item?.address} building_id={building_id}/>
                </li>
            }) : <li className={'no'}>Пусто</li>}
        </ul>
    )
}

function Building({id, name, address}) {
    const [collapsed, setCollapsed] = useState(false)
    function handler(event) {
        event.stopPropagation()
        setCollapsed(!collapsed)
    }
    return <div className={'zhk'} onClick={handler}>
        <div className={'item'}>
            {collapsed ? <img src={chevron_down}/> : <img src={chevron_up}/>}
            <span>{name}</span>
        </div>
        {collapsed ? <Elevators building_id={id}/> : <></>}
    </div>
}

function Buildings({complex_id}) {
    const {data, loading, error} = useFetch(`http://62.84.99.184:80/api/complexes/${complex_id}/buildings`)
    if (loading) {
        return <>Загрузка</>
    }
    return (
        <ul>
            {data?.buildings ? data?.buildings?.map(item => {
                return <li><Building id={item?.id} name={item?.name} address={item?.address}/></li>
            }) : <li className={'no'}>Пусто</li>}
        </ul>
    )
}

function Complex({id, name}) {
    const [collapsed, setCollapsed] = useState(false)

    function handler(event) {
        event.stopPropagation()
        setCollapsed(!collapsed)
    }

    return <div className={'zhk'} onClick={handler}>
        <div className={'item'}>
            {collapsed ? <img src={chevron_down}/> : <img src={chevron_up}/>}
            <span>{name}</span>
        </div>
        {collapsed ? <Buildings complex_id={id}/> : <></>}
    </div>
}


function Menu() {
    const {data, loading, error} = useFetch(`http://62.84.99.184:80/api/complexes`)
    return (
        <nav className={'card'}>
            <div className={'card-header'}>
                Объекты
            </div>
            <div className="menu_body">
                <ul>
                    {loading && <>Загрузка</>}
                    {data?.complexes?.map(item => {
                        return (
                            <li><Complex id={item?.id} name={item?.name}/></li>
                        )
                    })}
                </ul>
            </div>

        </nav>
    )
}

function Header() {
    return <div className={'header'}>
        <h1>
            <img src={logo}/>
            <span>LiftAssistant</span>
        </h1>
    </div>
}


function Help() {
    return (
        <ol>
            <li>Выберите Комплекс зданий, дом, лифт и соотвествующий экран в меню слева</li>
            <li>После выбора вы попадете в интерфейс редактирования экрана </li>
            <li>В данном интерфейсе вы можете изменять размеры виджетов, перемещать их</li>
            <li>Изменения сохраняются автоматически</li>
            <li>Для просмотра измененного экрана, нажмите кнопку play, находящуюся в верхнем правом углу карточки</li>
        </ol>
    )
}

function PlayButton({screen_id}) {
    const navigate = useNavigate()
    useEffect(() => {
        console.log(screen_id)
    }, [screen_id])
    function handleClick() {
        navigate(`/screen/${screen_id}`)
    }
    return (
        <div className='button_play' onClick={handleClick}>
            <img src={play}/>
        </div>
    )
}

function EditorSpawn() {
    const params = useParams()
    const {
        data,
        loading,
        error
    } = useFetch(`http://62.84.99.184:80/api/screen_widgets/${params?.screen_id}`)

    const [random, setRandom] = useState(0)
    useEffect(() => {
        setRandom(Math.random())
    }, [params])

    return (
        <div className="content" key={random}>
            <div className="card templates">
                <div className="card-header  d-flex">
                    Шаблоны
                    <PlayButton screen_id={params?.screen_id}/>
                </div>
                <div className="card-body">
                    {data?.screen_widgets?.length &&
                        <Editor screen_id={params?.screen_id} initial={data?.screen_widgets}/>}
                </div>
            </div>
        </div>
    )
}

function Admin() {

    useEffect(() => {
        document.title = 'My Admin Page';
    }, []);

    const params = useParams()

    return (
        <>
            <div className={'admin'}>
                <Header/>
                <Menu/>

                {params?.screen_id ? <EditorSpawn/> : <Help/>}
            </div>
        </>
    )
}


export default Admin
