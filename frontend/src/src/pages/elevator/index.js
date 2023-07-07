import './index.scss'
import React, {useState, useEffect} from "react";
import {Responsive, WidthProvider} from "react-grid-layout";

// import "/node_modules/react-grid-layout/css/styles.css";
// import "/node_modules/react-resizable/css/styles.css";
import USDWidget from "../../widgets/usd";
import useFetch from "../../useFetch";
import WeatherWidget from "../../widgets/weather";
import TimeWidget from "../../widgets/time";
import NewsWidget from "../../widgets/news";
import TransportWidget from "../../widgets/transport";
import TrafficJamsWidget from "../../widgets/traffic_jams";
import PlaceHolder from "../../widgets/place_holder";
import ButtonsGroup from "../../widgets/buttons_group";
import {useParams} from "react-router-dom";

const ResponsiveReactGridLayout = WidthProvider(Responsive);

function Editor({layout}) {
    if (!layout) {
        return <></>
    }

    return (
        <div className={'client_screen'}>
            <div
                className="droppable-element"
                draggable={true}
                unselectable="on"
                // this is a hack for firefox
                // Firefox requires some kind of initialization
                // which we can do by adding this attribute
                // @see https://bugzilla.mozilla.org/show_bug.cgi?id=568313
                onDragStart={e => e.dataTransfer.setData("text/plain", "")}
            >
            </div>
            <ResponsiveReactGridLayout
                {...this?.props}
                rowHeight={50}
                cols={{lg: 12, md: 12, sm: 12, xs: 12, xxs: 12}}
                layout={layout}

                // WidthProvider option
                measureBeforeMount={false}
                // I like to have it animate on mount. If you don't, delete `useCSSTransforms` (it's default `true`)
                // and set `measureBeforeMount={true}`.
                // useCSSTransforms={mounted}
                // compactType={'vertical'}
                compactType={false}
                onLayoutChange={() => {
                }}
                margin={[10, 10]}
                isResizable={false}
                isBounded={false}
                // allowOverlap={true}  // наложение друг на друга
                preventCollision={true}
                isDroppable={false}
                // onResize={FunconResize}
                droppingItem={{i: "xx", h: 50, w: 250}}
            >
                {layout.map((itm, i) => (
                    <div key={itm.i} data-grid={itm} className="block">
                        {
                            (itm.i === '63baeddd-2a07-4f71-aa19-62ecbae26429') && <USDWidget/>
                        }
                        {
                            (itm.i === 'd30bb91e-6718-4380-a196-9b791b26280d') && <WeatherWidget/>
                        }
                        {
                            (itm.i === 'd6e2f387-a6ea-471b-96c3-d46a0e7c796d') && <TimeWidget/>
                        }
                        {
                            (itm.i === 'b71bef49-574e-4354-867c-ca77794172be') && <NewsWidget/>
                        }
                        {
                            (itm.i === 'e6b16a02-3d14-4185-b02b-ef1c3035f159') && <TransportWidget/>
                        }
                        {
                            (itm.i === '070f62e1-dad3-454c-b89f-78df02df1039') && <TrafficJamsWidget/>
                        }
                        {
                            (itm.i === '61493b97-7d24-4957-9d0a-3548f456374f') && <PlaceHolder/>
                        }
                        {
                            (itm.i === 'e953c6b2-ce4d-42a1-b1b0-7a264172b1a2') && <PlaceHolder/>
                        }
                        {
                            (itm.i === '7e551a5b-ff79-4c4e-81c7-2697478d6b54') && <ButtonsGroup/>
                        }
                    </div>
                ))}
            </ResponsiveReactGridLayout>
        </div>
    );
}


function Elevator() {
    const params = useParams()

    useEffect(() => {
        document.title = 'Welcome';
    }, []);

    const {
        data,
        loading,
        error
    } = useFetch(`http://62.84.99.184:80/api/screen_widgets/${params?.screen_id}`)

    const [items, setItems] = useState([])

    useEffect(() => {
        if (data?.screen_widgets?.length) {
            data?.screen_widgets?.map(item => {
                item['static'] = true
                setItems(_old => [..._old, item])
                return 1
            })
            setItems(data?.screen_widgets)
        }
    }, [data])

    return (
        <div className={'elevator_client'}>
            {items && <Editor layout={items}/>}
        </div>
    )
}


export default Elevator
