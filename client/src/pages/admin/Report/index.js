import Table from "react-bootstrap/Table";
import { useEffect, useState } from "react";
import { Accordion, Badge } from "react-bootstrap";
import { useDispatch, useSelector } from "react-redux";

import "./style.css";
import Card from "../../../components/Shares/Card";
import SubCard from "../../../components/Shares/Card/SubCard";
import EmptyListNoti from "../../../components/Shares/EmptyListNoti";
import PaginationCustom from "../../../components/Shares/Pagination";
import Action from "../../../components/Shares/Action";
import { LIMIT } from "../../../const/const";
import Detail from "../../../components/Shares/Action/Detail";
import { topicSelector } from "../../../store/selectors";
import { fetchTopics } from "../../common/Topic/TopicSlice";
import ConfirmStatus from "../../../components/Shares/Action/ConfirmStatus";
import Cancel from "../../../components/Shares/Action/Cancel";

const Report = () => {
  const [pagi, setPagi] = useState(1);
  const dispatch = useDispatch();
  useEffect(() => {
    dispatch(
      fetchTopics({
        search: "",
        departmentID: 0,
        facultyID: 0,
        status: 0,
        timeStart: 0,
        timeEnd: 0,
        limit: LIMIT,
        offset: (pagi - 1) * LIMIT,
      })
    );
    // eslint-disable-next-line
  }, [dispatch, pagi]);

  const topic = useSelector(topicSelector);

  const listTopic = topic.topics;
  let total = topic.total;
  return (
    <Card title={"Nghiệm thu đề tài"}>
      <Accordion defaultActiveKey={"0"}>
        <Accordion.Item eventKey="0">
          <Accordion.Header className="add-topic-header">
            Đại hội nghiên cứu khoa học lần thứ 36 HUMG
          </Accordion.Header>
          <Accordion.Body>
          <SubCard title={"Danh sách các đề tài báo cáo cấp trường"}>
              {total !== 0 ? (
                <EmptyListNoti title={"Không có đề tài nào!"} />
              ) : (
                <div>
                </div>
              )}
            </SubCard>
            <SubCard title={"Danh sách các đề tài đã báo cáo tại các tiểu ban"}>
              {total === 0 ? (
                <EmptyListNoti title={"Không có đề tài nào!"} />
              ) : (
                <div>
                  <Table bordered hover size="sm" className="topic-table">
                    <thead>
                      <tr>
                        <th>STT</th>
                        <th style={{ width: "150px" }}>Giảng viên</th>
                        <th>Liên hệ</th>
                        <th>Đề tài</th>
                        <th style={{ width: "100px" }}>Khoa</th>
                        <th style={{ width: "100px" }}>Tiểu ban</th>
                        <th style={{ width: "100px" }}>Kết quả 1</th>
                        <th>Thao tác</th>
                      </tr>
                    </thead>
                    <tbody>
                      {listTopic.map((item, index) => {
                        return (
                          <tr key={index}>
                            <td>{(pagi - 1) * LIMIT + index + 1}</td>
                            <td>
                              {item.lectureInfo.degree +
                                " " +
                                item.lectureInfo.name}
                            </td>
                            <td style={{ textAlign: "center" }}>
                              {item.lectureInfo.email}
                              <br />
                              {item.lectureInfo.phone}
                            </td>
                            <td>{item.name}</td>
                            <td style={{ textAlign: "center" }}>
                              {"Công nghệ thông tin"}
                            </td>
                            <td style={{ textAlign: "center" }}>
                              {"Tin học kinh tế"}
                            </td>
                            <td style={{ textAlign: "center" }}>
                              <Badge bg="danger" style={{ width: "50px" }}>
                                {Math.floor(Math.random() * 10) + 20 + "đ"}
                              </Badge>
                            </td>
                           
                            <td>
                              <Action
                                todo={[
                                  <Detail name={"Xem chi tiết"} topic={item} />,
                                  <ConfirmStatus name={"Duyệt"} />,
                                  <Cancel name={"Hủy"} />,
                                ]}
                              />
                            </td>
                          </tr>
                        );
                      })}
                    </tbody>
                  </Table>

                  <PaginationCustom
                    setPagi={setPagi}
                    currentPage={pagi}
                    total={total}
                    limit={LIMIT}
                  />
                </div>
              )}
            </SubCard>
          </Accordion.Body>
        </Accordion.Item>
      </Accordion>
    </Card>
  );
};

export default Report;