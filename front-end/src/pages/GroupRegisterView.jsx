import '@/pages/GroupRegisterView.scss';
import { Badge, Button, Col, Form, Row } from 'react-bootstrap';
import { useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';

export default function GroupRegisterView(props) {
  const navigate = useNavigate();
  const [formData, setFormData] = useState({
    name: '',
    emailToInvite: '',
  });
  const [invitedEmails, setInvitedEmails] = useState([]);

  function handleFormDataChanged(evt) {
    const { name, value } = evt.target;
    setFormData(formData => ({
      ...formData,
      [name]: value
    }));
  }

  function addInvitedEmail() {
    if (invitedEmails.includes(formData.emailToInvite))
      return;
    setInvitedEmails(invitedEmails => invitedEmails.concat(formData.emailToInvite));
    setFormData(formData => ({
      ...formData,
      emailToInvite: '',
    }));
  }

  function removeInvitedEmail(invitedEmail) {
    const index = invitedEmails.indexOf(invitedEmail);
    if (index < 0)
      return;
    const _invitedEmails = [...invitedEmails];
    _invitedEmails.splice(index, 1);
    setInvitedEmails(_invitedEmails);
  }

  async function submit(evt) {
    evt.preventDefault();
    try {
      await axios.post('/v1/group', {
        name: formData.name,
        users: [props.userInfo.email].concat(invitedEmails),
        modUserId: props.userInfo.id,
        paymentMethods: []
      });
      navigate(-1);
    } catch (e) {
      console.log(e);
    }
  }

  return (
    <article className="abs-group-register">
      <Form
        className="register-form"
        onSubmit={submit}
      >
        <Form.Group className="register-row">
          <Form.Control
            name="name"
            type="string"
            value={formData.name}
            placeholder="그룹 이름"
            onChange={handleFormDataChanged}
            required
          />
        </Form.Group>
        <Form.Group className="register-row-group">
          <Form.Label>그룹 구성원 추가하기</Form.Label>
          <Row className="register-row">
            <Col
              as={Form.Control}
              name="emailToInvite"
              value={formData.emailToInvite}
              placeholder="초대할 이메일"
              onChange={handleFormDataChanged}
            />
            <Col
              as={Button}
              className="col-3"
              variant="outline-primary"
              // TODO 이메일 형식에 맞지 않으면 disable
              // TODO 이미 추가한 이메일이면 disable
              onClick={addInvitedEmail}
            >
              추가
            </Col>
          </Row>
          <ul className="invited-emails">
            {
              invitedEmails.map((email, i) => {
                let bg, text;
                if (i % 2) {
                  bg = 'light';
                  text = 'dark';
                } else
                  bg = 'info';

                return (
                  <li
                    key={email}
                    className="li__email"
                  >
                    <Badge
                      size="sm"
                      bg={bg}
                      text={text}
                    >
                      {email}
                    </Badge>
                    <Button
                      size="xs"
                      variant="soft-clear"
                      onClick={() => removeInvitedEmail(email)}
                    >x</Button>
                  </li>
                );
              })
            }
          </ul>
        </Form.Group>
        <Button
          className="w-100"
          type="submit"
        >
          등록
        </Button>
      </Form>
    </article>
  );
}
