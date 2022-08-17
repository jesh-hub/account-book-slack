import { useNavigate } from 'react-router-dom';

export default function useRouterNavigateWith() {
  const navigate = useNavigate();
  return (path, state) => {
    navigate(path, { state });
  };
}
