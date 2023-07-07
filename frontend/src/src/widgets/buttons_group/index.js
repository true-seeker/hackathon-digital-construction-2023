import './index.scss'


function ButtonsGroup() {
    const floors = [13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, -1];
    const start_floor = 12;
    const end_flor = 5;
    const current_flor = 7;

    return (
        <div className='buttons_groups'>
            <div className="buttons">
                {
                    floors?.map(item => {
                        let class_name = 'circle'
                        if (item === start_floor) {

                        }
                        if (item === end_flor) {
                            class_name += ' end'
                        }
                        if (item === current_flor) {
                            class_name += ' current'
                        }
                        return (
                            <div className={class_name}>{item}</div>
                        )
                    })
                }
            </div>
            <div className="progress">
                <div className="end" style={{height: '47%'}}></div>
            </div>
        </div>
    )
}


export default ButtonsGroup
