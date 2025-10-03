-- musicroom 사용자에게 모든 권한 부여
GRANT ALL PRIVILEGES ON *.* TO 'musicroom'@'%';

-- 권한 적용
FLUSH PRIVILEGES;